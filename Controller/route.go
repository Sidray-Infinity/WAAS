package Controller

import (
	"net/http"
)

func Route() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/get", Get)           // Test API
	mux.HandleFunc("/register", Register) // Use restful resources
	mux.HandleFunc("/addWallet", AddWallet)
	mux.HandleFunc("/credit", Credit)
	mux.HandleFunc("/debit", Debit)
	mux.HandleFunc("/block", Block)
	mux.HandleFunc("/unblock", UnBlock)
	mux.HandleFunc("/generateCSV", GenerateCSV)
	return mux
}
