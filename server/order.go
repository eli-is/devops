// B"H

/*
https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
*/
package server

import (
	"encoding/json"
	"fmt"
	"go-http-server/structs"
	"log"
	"net/http"
)

func Order(w http.ResponseWriter, r *http.Request) {
	var order structs.Order
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Data is incorrect: %+v\n", http.StatusBadRequest)
		// http.Error(w, "Data is incorrect", http.StatusBadRequest)
		return
	}
	// Return the Ingress structto the sender.
	log.Printf("Got Order: %+v ", order.Size)
	fmt.Fprintf(w, "Got Order: %+v", order.Size)

}
