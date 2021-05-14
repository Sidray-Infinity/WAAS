package Docs

import (
	entity "waas/Model/entity"
)

// type userRegisterReponse struct{}
type userRegisterRequest struct {
	UserName     string `json:"user_name"`
	Password     string `json:"password"`
	AadharNumber int    `json:"aadhar_number"`
}

// swagger:route POST /user User userRegister
// Api used to register a user.
// responses:
//   201: userRegisterResponse

// The only response is the status code.
// swagger:response userRegisterResponse
type userRegisterResponseWrapper struct{}

// swagger:parameters userRegister
type userRegisterParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body userRegisterRequest
}

// swagger:route GET /user/{Id} User userFetch
// API used to fetch user details.
// responses:
//   200: userFetchResponse

// The user details is returned in JSON format.
// swagger:response userFetchResponse
type userFetchResponseWrapper struct {
	// in:body
	Body entity.User
}

// swagger:parameters userFetch
type userFetchParamsWrapper struct {
	// ID of the user to be fetched.
	// in:path
	Id int
}
