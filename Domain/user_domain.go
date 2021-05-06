package Domain

import (
	"waas/Model/Impl"
	entity "waas/Model/entity"
)

func Get() []entity.User {
	return Impl.Get()
}

func Register(newUser entity.User) {
	Impl.Register(newUser)
}
