package Controller

import (
	"log"
	"net/http"
	"strconv"
	"waas/Domain"
)

func AddWallet(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.URL.Query()["user_id"]
	if !ok || len(userId[0]) < 1 {
		log.Println("Url Param 'user_id' is missing")
		return
	}
	userIdNum, err := strconv.Atoi(userId[0])
	if err != nil {
		log.Println("User ID is not numeric!", err)
		return
	}

	go Domain.AddWallet(userIdNum)

}

func Credit(w http.ResponseWriter, r *http.Request) {
	walletId, ok := r.URL.Query()["wallet_id"]
	if !ok || len(walletId[0]) < 1 {
		log.Println("Url Param 'walletId' is missing")
		return
	}
	walletIdNum, err := strconv.Atoi(walletId[0])
	if err != nil {
		log.Println("walletId is not numeric!", err)
		return
	}

	amount, ok := r.URL.Query()["amount"]
	if !ok || len(amount[0]) < 1 {
		log.Println("Url Param 'amount' is missing")
		return
	}
	amountNum, err := strconv.ParseFloat(amount[0], 64)
	if err != nil {
		log.Println("User ID is not numeric!", err)
		return
	}

	go Domain.Credit(walletIdNum, amountNum)

}

func Debit(w http.ResponseWriter, r *http.Request) {
	walletId, ok := r.URL.Query()["wallet_id"]
	if !ok || len(walletId[0]) < 1 {
		log.Println("Url Param 'walletId' is missing")
		return
	}
	walletIdNum, err := strconv.Atoi(walletId[0])
	if err != nil {
		log.Println("walletId is not numeric!", err)
	}

	amount, ok := r.URL.Query()["amount"]
	if !ok || len(amount[0]) < 1 {
		log.Println("Url Param 'amount' is missing")
		return
	}
	amountNum, err := strconv.ParseFloat(amount[0], 64)
	if err != nil {
		log.Println("User ID is not numeric!", err)
	}

	go Domain.Debit(walletIdNum, amountNum)

}

func Block(w http.ResponseWriter, r *http.Request) {
	walletId, ok := r.URL.Query()["wallet_id"]
	if !ok || len(walletId[0]) < 1 {
		log.Println("Url Param 'walletId' is missing")
		return
	}
	walletIdNum, err := strconv.Atoi(walletId[0])
	if err != nil {
		log.Println("walletId is not numeric!", err)
	}

	go Domain.Block(walletIdNum)
}

func UnBlock(w http.ResponseWriter, r *http.Request) {
	walletId, ok := r.URL.Query()["wallet_id"]
	if !ok || len(walletId[0]) < 1 {
		log.Println("Url Param 'walletId' is missing")
		return
	}
	walletIdNum, err := strconv.Atoi(walletId[0])
	if err != nil {
		log.Println("walletId is not numeric!", err)
	}

	go Domain.UnBlock(walletIdNum)
}
