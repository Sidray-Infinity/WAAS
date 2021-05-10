package Domain

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"waas/Model/Impl"
	entity "waas/Model/entity"

	"github.com/gorilla/mux"
)

func GetUser(rw http.ResponseWriter, r *http.Request) (*entity.User, error) {
	userId, _ := strconv.Atoi(mux.Vars(r)["id"])
	return Impl.GetUser(userId)
}

func RegisterUser(rw http.ResponseWriter, r *http.Request) error {
	newUser := entity.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Println("Cannot decode JSON", err)
		return err
	}
	return Impl.RegisterUser(newUser)
}
