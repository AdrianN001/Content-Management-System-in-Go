package main

import (
	"webserver/app"
)

func main() {
	server := app.App{Port: 0, Socket: nil}
	server.Init_webserver("localhost", 6969, app.Proccess_Connection)
}
