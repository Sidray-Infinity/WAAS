package Domain

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"waas/Model"
	"waas/Model/Impl"
	entity "waas/Model/entity"

	"github.com/gorilla/mux"
)

type UserDomainImpl struct {
	userModel Model.UserModel
}

func (u *UserDomainImpl) GetUser(rw http.ResponseWriter, r *http.Request) (*entity.User, error) {
	u.userModel = &Impl.UserModelImpl{}
	userId, _ := strconv.Atoi(mux.Vars(r)["id"])
	return u.userModel.GetUser(userId)
}

func (u *UserDomainImpl) RegisterUser(rw http.ResponseWriter, r *http.Request) error {
	u.userModel = &Impl.UserModelImpl{}
	newUser := entity.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Println("Cannot decode JSON", err)
		return err
	}
	return u.userModel.RegisterUser(newUser)
}
