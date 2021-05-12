package Domain

import (
	"net/http"
	entity "waas/Model/entity"
)

type WalletDomain interface {
	GetWallet(rw http.ResponseWriter, r *http.Request) (*entity.Wallet, error)
	RegisterWallet(rw http.ResponseWriter, r *http.Request) error
	GetBalance(rw http.ResponseWriter, r *http.Request) (*float64, error)
	GetStatus(rw http.ResponseWriter, r *http.Request) (*bool, error)
	WalletBalance(rw http.ResponseWriter, r *http.Request) (float64, int, error)
	WalletStatus(rw http.ResponseWriter, r *http.Request) error
}
