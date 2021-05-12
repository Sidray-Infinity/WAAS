package Controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Route() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/user/{id:[0-9]+}", userHandler).Methods("GET")
	mux.HandleFunc("/user", userHandler).Methods("POST")

	mux.HandleFunc("/wallet/{id:[0-9]+}", walletHandler).Methods("GET")
	mux.HandleFunc("/wallet", walletHandler).Methods("POST")
	mux.HandleFunc("/wallet/balance/{id:[0-9]+}", balanceHandler).Methods("PATCH", "GET")
	mux.HandleFunc("/wallet/status/{id:[0-9]+}", statusHandler).Methods("PATCH", "GET")
	mux.HandleFunc("/generateCSV", GenerateCSV)

	return mux
}
