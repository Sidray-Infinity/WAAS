package Controller

import (
	"encoding/json"
	"log"
	"net/http"
	"waas/Domain"
	DomainImpl "waas/Domain/Impl"
	"waas/Model/view"
)

type WalletHandler struct {
	walletDomain Domain.WalletDomain
}

func (w *WalletHandler) walletHandler(rw http.ResponseWriter, r *http.Request) {
	w.walletDomain = &DomainImpl.WalletDomainImpl{}
	if r.Method == http.MethodGet {
		// Fetch a wallet

		wallet, err := w.walletDomain.GetWallet(rw, r)
		if err != nil {
			http.Error(rw, "Cannot fetch wallet", http.StatusInternalServerError) // Error code should be decided at runtime
			return
		}
		json.NewEncoder(rw).Encode(wallet)
	} else if r.Method == http.MethodPost {
		// Add a new wallet

		err := w.walletDomain.RegisterWallet(rw, r)
		if err != nil {
			http.Error(rw, "Cannot Regsiter wallet", http.StatusInternalServerError)
			return
		}
		log.Println("Wallet Created")
		rw.WriteHeader(http.StatusCreated)
	}
}

func (w *WalletHandler) balanceHandler(rw http.ResponseWriter, r *http.Request) {
	w.walletDomain = &DomainImpl.WalletDomainImpl{}
	if r.Method == http.MethodPatch {
		updatedBalance, txnId, err := w.walletDomain.WalletBalance(rw, r)
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
		balance, err := w.walletDomain.GetBalance(rw, r)
		if err != nil {
			http.Error(rw, "Could not fetch wallet balance", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(map[string]float64{"balance": *balance})
	}
}

func (w *WalletHandler) statusHandler(rw http.ResponseWriter, r *http.Request) {
	w.walletDomain = &DomainImpl.WalletDomainImpl{}
	if r.Method == http.MethodPatch {
		err := w.walletDomain.WalletStatus(rw, r)
		if err != nil {
			http.Error(rw, "Could not update wallet balance", http.StatusInternalServerError)
			return
		}
		log.Println("Wallet status updated")
		rw.WriteHeader(http.StatusNoContent)
	} else if r.Method == http.MethodGet {
		status, err := w.walletDomain.GetStatus(rw, r)
		if err != nil {
			http.Error(rw, "Could not fetch wallet status", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(map[string]bool{"status": *status})
	}
}
