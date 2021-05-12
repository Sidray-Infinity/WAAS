package Controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Route() http.Handler {

	w := &WalletHandler{}
	u := &UserHandler{}
	// t := &TransactionHandler{}

	mux := mux.NewRouter()

	mux.HandleFunc("/user/{id:[0-9]+}", u.userHandler).Methods("GET")
	mux.HandleFunc("/user", u.userHandler).Methods("POST")

	mux.HandleFunc("/wallet/{id:[0-9]+}", w.walletHandler).Methods("GET")
	mux.HandleFunc("/wallet", w.walletHandler).Methods("POST")
	mux.HandleFunc("/wallet/balance/{id:[0-9]+}", w.balanceHandler).Methods("PATCH", "GET")
	mux.HandleFunc("/wallet/status/{id:[0-9]+}", w.statusHandler).Methods("PATCH", "GET")
	// mux.HandleFunc("/generateCSV", GenerateCSV)

	return mux
}
