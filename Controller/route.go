package Controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Route() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/user/{id:[0-9]+}", user).Methods("GET")
	mux.HandleFunc("/user", user).Methods("POST")
	mux.HandleFunc("/wallet/{id:[0-9]+}", wallet).Methods("GET")
	mux.HandleFunc("/wallet", wallet).Methods("POST")
	mux.HandleFunc("/wallet/balance/{id:[0-9]+}", walletBalance).Methods("PATCH")
	mux.HandleFunc("/wallet/status/{id:[0-9]+}", walletStatus).Methods("PATCH")

	return mux
}
