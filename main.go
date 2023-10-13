package main

import (
	"fmt"
	"os"
	"strconv"
	"webserver/app"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 2{ 

		fmt.Println("Invalid Arguments")
		fmt.Println("Usage: \n\t ./main {address} {port}")
		return 
	}
	host := arguments[0]
	port, err := strconv.Atoi(arguments[1])
	if err != nil{
		fmt.Printf("Invalid port was given. %s \n", arguments[1])
		return
	}

	server := app.App{Port: 0, Socket: nil}
	server.Init_webserver(host, port, app.Proccess_Connection)
}
