package action

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"waas/manager"
	"waas/model"
)

var err error
var ok bool

func Register(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	userName, ok := r.URL.Query()["user_name"]
	if !ok || len(userName[0]) < 1 {
		log.Println("Url Param 'user_name' is missing")
		return
	}

	password, ok := r.URL.Query()["password"]
	if !ok || len(password[0]) < 1 {
		log.Println("Url Param 'password' is missing")
		return
	}

	aadharNum, ok := r.URL.Query()["aadhar_number"]
	if !ok || len(aadharNum[0]) < 1 {
		log.Println("Url Param 'aadhar_number' is missing")
		return
	}
	aadharNumber, err := strconv.Atoi(aadharNum[0])
	if err != nil {
		log.Println("Aadhar Number is not numeric!", err)
	}

	newUser.UserName = userName[0]
	newUser.Password = password[0]
	newUser.AadharNumber = aadharNumber

	manager.Register(newUser)
}

func Get(w http.ResponseWriter, r *http.Request) {
	users := manager.Get()
	jsonInfo, _ := json.Marshal(users)
	log.Printf("GET response: %s\n", string(jsonInfo))
	w.Write([]byte(string(jsonInfo)))
}

func AddWallet(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.URL.Query()["user_id"]
	if !ok || len(userId[0]) < 1 {
		log.Println("Url Param 'user_id' is missing")
		return
	}
	userIdNum, err := strconv.Atoi(userId[0])
	if err != nil {
		log.Println("User ID is not numeric!", err)
	}

	manager.AddWallet(userIdNum)

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

	manager.Credit(walletIdNum, amountNum)

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

	manager.Debit(walletIdNum, amountNum)

}
