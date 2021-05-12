package Domain

import (
	"net/http"
	entity "waas/Model/entity"
)

type UserDomain interface {
	GetUser(rw http.ResponseWriter, r *http.Request) (*entity.User, error)
	RegisterUser(rw http.ResponseWriter, r *http.Request) error
}
