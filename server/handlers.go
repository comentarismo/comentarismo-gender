package server

import (
	"net/http"
	"log"
	"encoding/json"
	"comentarismo-gender/gender"
)

func ReportGenderHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.

	lang := req.URL.Query().Get("lang")
	//validate inputs
	if lang == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing lang"})
		w.Write(jsonBytes)
		return
	}

	gen := req.URL.Query().Get("gender")
	if gen == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing gender argument"})
		w.Write(jsonBytes)
		return
	}

	g := gender.Gender(gen)

	//log.Println("lang , ", lang)
	if lang != "pt" && lang != "en" && lang != "fr" && lang != "es" && lang != "it" && lang != "hr" && lang != "ru" {
		errMsg := "Error: SentimentHandler Language " + lang + " not yet supported, use lang={en|pt|es|it|fr|hr|ru} eg lang=en"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	text := req.Form["text"]
	reply := gender.GenderReport{}

	//validate inputs
	if len(text) == 0 {
		errMsg := "Error: ReportSpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		reply.Code = 404
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("ReportSpamHandler, -->  ", text)
	gender.Train(g, text[0], lang)
	reply.Code = 200

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		reply.Code = 404
		errMsg := "Error: ReportSpamHandler Marshal"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func RevokeGenderHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	lang := req.URL.Query().Get("lang")

	//validate inputs
	if lang == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing lang"})
		w.Write(jsonBytes)
		return
	}

	gen := req.URL.Query().Get("gender")
	if gen == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing gender argument"})
		w.Write(jsonBytes)
		return
	}

	g := gender.Gender(gen)

	//log.Println("lang , ", lang)
	if lang != "pt" && lang != "en" && lang != "fr" && lang != "es" && lang != "it" && lang != "hr" && lang != "ru" {
		errMsg := "Error: SentimentHandler Language " + lang + " not yet supported, use lang={en|pt|es|it|fr|hr|ru} eg lang=en"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	text := req.Form["text"]
	reply := gender.GenderReport{}

	//validate inputs
	if len(text) == 0 {
		errMsg := "Error: RevokeSpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		reply.Code = 404
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("RevokeSpamHandler, -->  ", text)

	gender.Untrain(g, text[0], lang)

	reply.Code = 200

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		reply.Code = 404
		errMsg := "Error: RevokeSpamHandler Marshal"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func GenderHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//log.Println(req.Form)

	lang := req.URL.Query().Get("lang")

	//validate inputs
	if lang == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing lang"})
		w.Write(jsonBytes)
		return
	}

	//log.Println("lang , ", lang)
	if lang != "pt" && lang != "en" && lang != "fr" && lang != "es" && lang != "it" && lang != "hr" && lang != "ru" {
		errMsg := "Error: GenderHandler Language " + lang + " not yet supported, use lang={en|pt|es|it|fr|hr|ru} eg lang=en"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	text := req.Form["text"]
	reply := gender.GenderReport{}

	//validate inputs
	if len(text) == 0 {
		reply.Code = 404
		errMsg := "Error: GenderHandler text not defined, use eg: name=John"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	class := gender.Classify(text[0], lang)
	reply.Code = 200
	if class == "bad" {
		reply.Gender = "male"
	} else {
		reply.Gender = "female"
	}

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		reply.Code = 404
		errMsg := "Error: GenderHandler Marshal"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
