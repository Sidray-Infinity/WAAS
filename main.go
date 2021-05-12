package main

import (
	"log"
	"net/http"
	"os"
	controller "waas/Controller"
	"waas/Model/Impl"

	"github.com/jasonlvhit/gocron"
)

func main() {

	// Cron job to generate CSV periodically
	gocron.Every(1).Day().At("09:00").Do(controller.GenerateCSV)
	go func() {
		<-gocron.Start()
	}()

	log.Println("Connecting to DB ...")
	Impl.ConnnectToDB()
	defer Impl.CloseDB()

	log.Println("Connecting to Redis Client ...")
	Impl.ConnectRedis()
	defer Impl.CloseRedis()

	log.Println("Server starting ...")
	mux := controller.Route()
	log.Fatal(http.ListenAndServe("127.0.0.1:"+os.Args[1], mux))
	// log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
	// TODO : Implement graceful shutdown
}
