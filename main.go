package main

import (
	"net/http"
	"waas/action"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", action.Register)
	mux.HandleFunc("/get", action.Get)
	mux.HandleFunc("/addWallet", action.AddWallet)
	mux.HandleFunc("/credit", action.Credit)
	mux.HandleFunc("/debit", action.Debit)
	// mux.HandleFunc("/block", action.Block)
	// mux.HandleFunc("/unblock", action.UnBlock)
	// mux.HandleFunc("/generateCSV", action.GenerateCSV)

	http.ListenAndServe(":8080", mux)
}
