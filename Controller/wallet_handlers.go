package Controller

import (
	"encoding/json"
	"log"
	"net/http"
	"waas/Domain"
	"waas/Model/view"
)

func walletHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Fetch a wallet

		wallet, err := Domain.GetWallet(rw, r)
		if err != nil {
			http.Error(rw, "Cannot fetch wallet", http.StatusInternalServerError) // Error code should be decided at runtime
		}
		json.NewEncoder(rw).Encode(wallet)
	} else if r.Method == http.MethodPost {
		// Add a new wallet

		err := Domain.RegisterWallet(rw, r)
		if err != nil {
			http.Error(rw, "Cannot Regsiter wallet", http.StatusInternalServerError)
			return
		}
		log.Println("Wallet Created")
		rw.WriteHeader(http.StatusCreated)
	} else {
		// Catch undefined methods

		http.Error(rw, "Method not implemented for wallet", http.StatusBadRequest)
	}
}

func balanceHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		updatedBalance, txnId, err := Domain.WalletBalance(rw, r)
		if err != nil {
			http.Error(rw, "Could not update wallet balance", http.StatusInternalServerError)
			return
		}

		var resp view.BalanceUpdateResp
		resp.TransactionId = txnId
		resp.UpdatedBalance = updatedBalance
		json.NewEncoder(rw).Encode(resp)

		log.Println("Wallet balance updated")

	} else {
		// Catch undefined methods

		http.Error(rw, "Method not implemented for walletBalance", http.StatusBadRequest)
	}
}

func statusHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		err := Domain.WalletStatus(rw, r)
		if err != nil {
			http.Error(rw, "Could not update wallet balance", http.StatusInternalServerError)
			return
		}
		log.Println("Wallet status updated")
		rw.WriteHeader(http.StatusNoContent)
	} else {
		// Catch undefined methods

		http.Error(rw, "Method not implemented for walletBalance", http.StatusBadRequest)
	}
}
