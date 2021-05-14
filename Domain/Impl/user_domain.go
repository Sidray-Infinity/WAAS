package Domain

import (
	"encoding/json"
	"errors"
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
	err = u.validate(newUser)
	if err != nil {
		return err
	}
	return u.userModel.RegisterUser(newUser)
}

func (u *UserDomainImpl) validate(newUser entity.User) error {
	if !u.userModel.ValidateUsername(newUser) {
		return errors.New("username invalid or taken")
	}
	if !u.userModel.ValidateKYC(newUser) {
		return errors.New("invalid Kyc")
	}
	return nil
}
