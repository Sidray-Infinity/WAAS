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

	} else if r.Method == http.MethodGet {
		balance, err := Domain.GetBalance(rw, r)
		if err != nil {
			http.Error(rw, "Could not fetch wallet balance", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(map[string]float64{"balance": *balance})
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
	} else if r.Method == http.MethodGet {
		status, err := Domain.GetStatus(rw, r)
		if err != nil {
			http.Error(rw, "Could not fetch wallet status", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(map[string]bool{"status": *status})
	}
}
