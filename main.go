package main

import (
	"log"
	"net/http"
	controller "waas/Controller"
)

func main() {
	mux := controller.Route()
	log.Println("Server starting ...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("Cannot start server:", err)
		return
	}

}
