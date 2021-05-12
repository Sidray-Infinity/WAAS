package Controller

import (
	"encoding/json"
	"net/http"
	Domain "waas/Domain"
	DomainImpl "waas/Domain/Impl"
)

type UserHandler struct {
	userDomain Domain.UserDomain
}

func (u *UserHandler) userHandler(rw http.ResponseWriter, r *http.Request) {
	u.userDomain = &DomainImpl.UserDomainImpl{}
	if r.Method == http.MethodGet {
		user, err := u.userDomain.GetUser(rw, r)
		if err != nil {
			http.Error(rw, "Cannot fetch user", http.StatusInternalServerError)
		}
		json.NewEncoder(rw).Encode(user)

	} else if r.Method == http.MethodPost {
		err := u.userDomain.RegisterUser(rw, r)
		if err != nil {
			http.Error(rw, "Cannot Regsiter user", http.StatusInternalServerError)
		}
		rw.WriteHeader(http.StatusCreated)
	} else {
		// catch all
		http.Error(rw, "Method not implemented for user", http.StatusBadRequest)
	}

}
