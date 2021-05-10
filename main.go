package main

import (
	"log"
	"net/http"
	controller "waas/Controller"
	"waas/Model/Impl"
)

func main() {
	mux := controller.Route()
	log.Println("Connecting to DB ...")
	Impl.ConnnectToDB()
	log.Println("Server starting ...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))

	// TODO : Implement graceful shutdown
}
