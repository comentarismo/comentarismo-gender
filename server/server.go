package server

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gorilla/pat"
	"log"
	"net/http"
	"os"
)

var (
	router *pat.Router
)

type WebError struct {
	Error string
}

func init() {
}

//NewServer return pointer to new created server object
func NewServer(Port string) *http.Server {
	router = InitRouting()
	return &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}
}

//StartServer start and listen @server
func StartServer(Port string) {
	log.Println("Starting server")
	s := NewServer(Port)
	fmt.Println("Server starting --> " + Port)
	err := gracehttp.Serve(
		s,
	)
	if err != nil {
		log.Fatalln("Error: %v", err)
		os.Exit(0)
	}

}

func InitRouting() *pat.Router {

	r := pat.New()

	/** bayes classifier spam **/
	r.Post("/gender", GenderHandler)
	r.Post("/revoke", RevokeGenderHandler)
	r.Post("/report", ReportGenderHandler)

	return r
}
