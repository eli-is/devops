// B"H
/*
https://stackoverflow.com/a/12655719/9202256 function that takes function.
https://www.alexedwards.net/blog/a-recap-of-request-handling
*/
package server

import (
	"fmt"
	"log"

	// "io/ioutil"
	"encoding/json"
	"io"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{
		Status: "up",
	})
}

type AnswerRequest func(w http.ResponseWriter, r *http.Request)

func DoGenericPost(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", reqBody)
	// fmt.Printf("%s\n", reqBody)
	w.Write([]byte("Received a POST request"))
}

func DoGenericGet(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.URL.Query() {
		fmt.Printf("%s %s\n", k, v)
	}
	w.Write([]byte("Received a GET request\n"))
}

func GenericHandler(path string, doPost AnswerRequest, doGet AnswerRequest) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			fmt.Printf("%s\n", path)
			http.NotFound(w, r)
			return
		}
		switch r.Method {
		case "GET":
			doGet(w, r)
		case "POST":
			doPost(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
		}
	}
	return http.HandlerFunc(fn)
}
