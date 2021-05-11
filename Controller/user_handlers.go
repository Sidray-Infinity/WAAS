package Controller

import (
	"encoding/json"
	"net/http"
	"waas/Domain"
)

func userHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user, err := Domain.GetUser(rw, r)
		if err != nil {
			http.Error(rw, "Cannot fetch user", http.StatusInternalServerError)
		}
		json.NewEncoder(rw).Encode(user)

	} else if r.Method == http.MethodPost {
		err := Domain.RegisterUser(rw, r)
		if err != nil {
			http.Error(rw, "Cannot Regsiter user", http.StatusInternalServerError)
		}
		rw.WriteHeader(http.StatusCreated)
	} else {
		// catch all
		http.Error(rw, "Method not implemented for user", http.StatusBadRequest)
	}

}
