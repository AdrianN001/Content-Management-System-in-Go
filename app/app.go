package app

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	SERVER_TYPE = "tcp"
)

type App struct {
	Port   int
	Socket *net.Listener

	// Init Method ( host, port )
}

func (this App) Init_webserver(host string, port int, process func(net.Conn)) {
	fmt.Println("Starting.....")

	server, err := net.Listen(SERVER_TYPE, fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		fmt.Println("Couldnt establish connection.")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	this.Socket = &server
	this.Port = port

	fmt.Printf(" Started listening on port %d \n", port)

	defer server.Close()

	for {
		connection, err := server.Accept()

		if err != nil {
			fmt.Print("Error accepting", err.Error())
		}
		go process(connection)

	}

}

func Proccess_Connection(Con net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := Con.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	var route string = strings.Split(string(buffer[:mLen]), " ")[1]

	fmt.Println("Received: ", string(buffer[:mLen]))

	_, err = Con.Write([]byte(fmt.Sprintf("This is your route %s", route)))
	Con.Close()
}
