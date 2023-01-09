package app

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	STATIC_FOLDER = "../static"
)

func create_embed(html_page, route string) string {
	f, err := os.Open(STATIC_FOLDER + route + html_page)
	if err != nil {
		fmt.Println("Cant open")
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var res []string

	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "<script") {
			var start int = strings.Index(scanner.Text(), "src=")
			var end int = strings.Index(scanner.Text(), ".js")

			script_path := scanner.Text()[start+5:end] + ".js"
			fmt.Println(script_path)

			js_file, err := ioutil.ReadFile(route + script_path)

			if err != nil {
				continue
			}
			var new_line string = fmt.Sprintf("<script defer>%s</script>", string(js_file))
			res = append(res, new_line)
		} else if strings.Contains(scanner.Text(), "stylesheet") {
			var start int = strings.Index(scanner.Text(), "href=")
			var end int = strings.Index(scanner.Text(), ".css")

			script_path := scanner.Text()[start+6:end] + ".css"
			fmt.Println(script_path)

			js_file, err := ioutil.ReadFile(route + script_path)

			if err != nil {
				continue
			}
			var new_line string = fmt.Sprintf("<style>\n%s</style>", string(js_file))
			res = append(res, new_line)

		} else {
			res = append(res, scanner.Text())
		}
	}
	return strings.Join(res, "\n")
}

func return_route(route string) (string, error) {
	// Returns a HTML page with embeded javascript and CSS

	files, err := ioutil.ReadDir(STATIC_FOLDER + route)

	if err != nil {
		fmt.Print("Route doesn't exist")
		return "", err
	}
	fmt.Print(files)

	if strings.HasSuffix(route, ".html") {
		return create_embed(route[strings.LastIndex(route, "/")+1:], route), nil
	}
	return create_embed("main.html", route), nil
}
