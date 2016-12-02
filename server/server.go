package server

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gorilla/pat"
	"log"
	"net/http"
	"os"
	"comentarismo-gender/gender"
)

var (
	router *pat.Router
)

type WebError struct {
	Error string
}

func init() {
	//
	if gender.LEARNGENDER == "true" {
		//will train with world know gender words
		var start = 1950
		var end = 2012
		log.Println("Will start server on learning mode")

		done := make(chan bool, end - start)
		for i := start; i <= end; i++ {
			targetFile := fmt.Sprintf("/gender/en/yob%d.txt", i)
			go gender.StartLanguageGender(targetFile, done)
		}
		go func() {
			for j := start; j <= end; j++ {
				targetFile := fmt.Sprintf("/gender/en/yob%d.txt", j)
				log.Println("Finished learning ", <-done, targetFile)
			}
		}()

	}
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
