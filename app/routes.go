package app

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const (
	STATIC_FOLDER = "static/"
	ASSET_FOLDER = "assets/"
)

func create_embed(html_page, route string) string {
	f, err := os.Open(STATIC_FOLDER + route + "/" + html_page)
	if err != nil {
		fmt.Printf("Cant open: %s", STATIC_FOLDER+"/"+route+"/"+html_page)
		return ""
	}
	defer f.Close()

	image_search_regex := regexp.MustCompile("src( |)=( |)\"[^\"]*\"")
	src_start_regex := regexp.MustCompile("src( |)=( |)")
	href_start_regex := regexp.MustCompile("href( |)=( |)")

	scanner := bufio.NewScanner(f)

	var res []string

	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "<script") {
			var start int = src_start_regex.FindStringIndex(scanner.Text())[0]
			var end int = strings.Index(scanner.Text(), ".js")

			script_path := scanner.Text()[start+5:end] + ".js"

			js_file, err := ioutil.ReadFile(STATIC_FOLDER + route + "/" + script_path)

			if err != nil {
				fmt.Printf("COULDN'T OPEN THE JS FILE: %s \n", STATIC_FOLDER+route+script_path)
				continue
			}
			var new_line string = fmt.Sprintf("<script defer>%s</script>", string(js_file))

			res = append(res, new_line)
		} else if strings.Contains(scanner.Text(), "stylesheet") {
			var start int = href_start_regex.FindStringIndex(scanner.Text())[0]
			var end int = strings.Index(scanner.Text(), ".css")

			script_path := scanner.Text()[start+6:end] + ".css"

			js_file, err := ioutil.ReadFile(STATIC_FOLDER + route + "/" + script_path)

			if err != nil {

				continue
			}
			var new_line string = fmt.Sprintf("<style>\n%s</style>", string(js_file))

			res = append(res, new_line)

		} else if strings.Contains(scanner.Text(), "<img"){

			var start int = src_start_regex.FindStringIndex(scanner.Text())[0]
			var end int = strings.Index(scanner.Text(), ".png")

			image_path := scanner.Text()[start+5:end] + ".png"

			file_in_folder := ASSET_FOLDER + image_path

			replaced_file_path := image_search_regex.ReplaceAllString(scanner.Text(), ( "src=\"" + file_in_folder + "\""))

			fmt.Println(image_path, file_in_folder, replaced_file_path)
			
			res = append(res, replaced_file_path)
		
		}else {
			res = append(res, scanner.Text())
		}
	}
	var html_string string = strings.Replace(strings.Join(res, ""), "\t", "", -1)
	//fmt.Println(html_string)
	return html_string
}

func return_route(route string) (string, error) {
	// Returns a HTML page with embeded javascript and CSS
	fmt.Println("ROUTE: " ,route)

	if strings.HasSuffix(route, ".html") {
		return create_embed(route[strings.LastIndex(route, "/")+1:], route[:strings.LastIndex(route, "/")]), nil
	}else if strings.HasSuffix(route, ".png") || strings.HasSuffix(route, ".jpg") || strings.HasSuffix(route, ".jpeg") || strings.HasSuffix(route, ".gif"){
		path_start_index := strings.Index(route, ASSET_FOLDER)

		var file_path_in_assets_folder string 
		if !strings.Contains(route, ASSET_FOLDER) || path_start_index == -1{
			// Asked from the CSS
			start_index := strings.LastIndex(route, "/")
			file_path_in_assets_folder = ASSET_FOLDER + route[start_index + 1:]
			
		}else{
			// Asked from the HTML

			file_path_in_assets_folder = route[path_start_index:]
		}
		fmt.Println(file_path_in_assets_folder, "@$RWF")


		buffer, err := ioutil.ReadFile(file_path_in_assets_folder)
		if err != nil{
			return "", err
		}
		
		return string(buffer), err
	}
	
	fmt.Println("NO HTML FILE WAS GIVEN. MAIN WILL BE RENDERED")
	return create_embed("main.html", route), nil
}
