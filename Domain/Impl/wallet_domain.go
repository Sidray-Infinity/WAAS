package Domain

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"waas/Model"
	"waas/Model/Impl"
	entity "waas/Model/entity"
	"waas/Model/view"

	"github.com/gorilla/mux"
)

type WalletDomainImpl struct {
	walletModel Model.WalletModel
}

func (w *WalletDomainImpl) GetWallet(rw http.ResponseWriter, r *http.Request) (*entity.Wallet, error) {
	w.walletModel = &Impl.WalletModelImpl{}
	userId, _ := strconv.Atoi(mux.Vars(r)["id"])
	return w.walletModel.GetWallet(userId)
}

func (w *WalletDomainImpl) RegisterWallet(rw http.ResponseWriter, r *http.Request) error {
	w.walletModel = &Impl.WalletModelImpl{}
	newWallet := &entity.Wallet{}
	err := json.NewDecoder(r.Body).Decode(&newWallet)
	if err != nil {
		log.Println("Cannot decode JSON", err)
		return err
	}
	return w.walletModel.RegisterWallet(newWallet)
}

func (w *WalletDomainImpl) GetBalance(rw http.ResponseWriter, r *http.Request) (*float64, error) {
	w.walletModel = &Impl.WalletModelImpl{}
	walletId, _ := strconv.Atoi(mux.Vars(r)["id"])
	return w.walletModel.GetBalance(walletId)
}

func (w *WalletDomainImpl) GetStatus(rw http.ResponseWriter, r *http.Request) (*bool, error) {
	w.walletModel = &Impl.WalletModelImpl{}
	walletId, _ := strconv.Atoi(mux.Vars(r)["id"])
	return w.walletModel.GetStatus(walletId)
}

func (w *WalletDomainImpl) WalletBalance(rw http.ResponseWriter, r *http.Request) (float64, int, error) {
	w.walletModel = &Impl.WalletModelImpl{}
	updateReq := &view.BalanceUpdate{}
	json.NewDecoder(r.Body).Decode(updateReq)
	walletId, _ := strconv.Atoi(mux.Vars(r)["id"])

	return w.walletModel.WalletBalance(updateReq, walletId)
}

func (w *WalletDomainImpl) WalletStatus(rw http.ResponseWriter, r *http.Request) error {
	w.walletModel = &Impl.WalletModelImpl{}
	updateReq := &view.StatusUpdate{}
	json.NewDecoder(r.Body).Decode(updateReq)
	walletId, _ := strconv.Atoi(mux.Vars(r)["id"])
	return w.walletModel.WalletStatus(updateReq, walletId)
}
