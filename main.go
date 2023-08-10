// B"H
package main

import (
	"fmt"
	"log"
	"net/http"

	"go-http-server/server"
	"go-http-server/utils"
)

func main() {

	conf := utils.ReadConfig()
	listenAt := fmt.Sprintf(":%d", conf.Port)

	handler := http.NewServeMux()

	healthHandle := server.GenericHandler("/health", server.Health, server.Health)
	genericHandle := server.GenericHandler("/", server.DoGenericPost, server.DoGenericGet)
	orderHandle := server.GenericHandler("/order", server.Order, server.DoGenericGet)

	fmt.Println("Server is starting at port: ", listenAt)

	handler.Handle("/health", healthHandle)
	handler.Handle("/", genericHandle)
	handler.Handle("/order", orderHandle)

	log.Fatal(http.ListenAndServe(listenAt, handler))
}
