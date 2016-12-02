package main

import (
	"comentarismo-gender/server"
	"os"
)

var Port = os.Getenv("PORT")

func main() {
	if Port == "" {
		Port = "3005"
	}
	server.StartServer(Port)
}
