package Domain

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"waas/Model/Impl"
	entity "waas/Model/entity"
	"waas/Model/view"

	"github.com/gorilla/mux"
)

func GetWallet(rw http.ResponseWriter, r *http.Request) (*entity.Wallet, error) {
	userId, _ := strconv.Atoi(mux.Vars(r)["id"])
	return Impl.GetWallet(userId)
}

func RegisterWallet(rw http.ResponseWriter, r *http.Request) error {
	newWallet := &entity.Wallet{}
	err := json.NewDecoder(r.Body).Decode(&newWallet)
	if err != nil {
		log.Println("Cannot decode JSON", err)
		return err
	}
	return Impl.RegisterWallet(newWallet)
}

func WalletBalance(rw http.ResponseWriter, r *http.Request) (float64, int, error) {
	updateReq := &view.BalanceUpdate{}
	json.NewDecoder(r.Body).Decode(updateReq)
	walletId, _ := strconv.Atoi(mux.Vars(r)["id"])
	return Impl.WalletBalance(updateReq, walletId)
}

func WalletStatus(rw http.ResponseWriter, r *http.Request) error {
	updateReq := &view.StatusUpdate{}
	json.NewDecoder(r.Body).Decode(updateReq)
	walletId, _ := strconv.Atoi(mux.Vars(r)["id"])
	return Impl.WalletStatus(updateReq, walletId)
}
