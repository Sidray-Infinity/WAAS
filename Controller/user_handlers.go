package Controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"waas/Domain"
	entity "waas/Model/entity"
)

func Get(w http.ResponseWriter, r *http.Request) {
	users := Domain.Get()
	jsonInfo, _ := json.Marshal(users)
	log.Printf("GET response: %s\n", string(jsonInfo))
	w.Write([]byte(string(jsonInfo)))
}

func Register(w http.ResponseWriter, r *http.Request) {
	var newUser entity.User
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
		return
	}

	newUser.UserName = userName[0]
	newUser.Password = password[0]
	newUser.AadharNumber = aadharNumber

	go Domain.Register(newUser)
}
