package main

import (
	"os"
	"comentarismo-gender/server"
)

var Port = os.Getenv("PORT")

func main() {
	if Port == "" {
		Port = "3004"
	}
	server.StartServer(Port)
}
