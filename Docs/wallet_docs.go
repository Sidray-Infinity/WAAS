package Docs

import (
	entity "waas/Model/entity"
	"waas/Model/view"
)

type walletRegisterRequest struct {
	UserId int `json:"user_id"`
}

type balanceFetchResponse struct {
	Balance int `json:"balance"`
}

type statusUpdateRequest struct {
	Status bool `json:"status"`
}

type statusFetchResponse struct {
	Status bool `json:"status"`
}

// swagger:route POST /wallet Wallet walletRegister
// Api used to register a wallet, for a given user.
// responses:
//   201: walletRegisterResponse

// The only response is the status code.
// swagger:response walletRegisterResponse
type walletRegisterResponseWrapper struct{}

// swagger:parameters walletRegister
type walletRegisterParamsWrapper struct {
	// UserId of the user for which wallet has to be created, is passed as input.
	// in:body
	Body walletRegisterRequest
}

// swagger:route GET /wallet/{Id} Wallet walletFetch
// Api used to fetch a wallet, based on given wallet id.
// responses:
//   200: walletFetchResponse

// The wallet record is returned in JSON format.
// swagger:response walletFetchResponse
type walletFetchResponseWrapper struct {
	// in:body
	Body entity.Wallet
}

// swagger:parameters walletFetch
type walletFetchParamsWrapper struct {
	// Wallet id is passed as parameter.
	// in:path
	Id int
}

// swagger:route PATCH /wallet/balance/{Id} Wallet walletBalanceUpdate
// Api used to update the balance of a wallet, based on given wallet id.
// responses:
//   200: walletBalanceUpdateResponse

// The transaction id, as well as updated balance of the wallet is returned.
// swagger:response walletBalanceUpdateResponse
type walletBalanceUpdateResponse struct {
	// in:body
	Body view.BalanceUpdateResp
}

// swagger:parameters walletBalanceUpdate
type walletBalanceUpdateParamsWrapper struct {
	// Wallet id is passed as parameter.
	// in:path
	Id int
}

// swagger:route GET /wallet/balance/{Id} Wallet walletBalanceFetch
// Api used to fetch the balance of a wallet, based on given wallet id.
// responses:
//   200: walletBalanceFetchResponse

// The current balance of the wallet is returned.
// swagger:response walletBalanceFetchResponse
type walletBalanceFetchResponse struct {
	// in:body
	Body balanceFetchResponse
}

// swagger:parameters walletBalanceFetch
type walletBalanceFetchParamsWrapper struct {
	// Wallet id is passed as parameter.
	// in:path
	Id int
}

// swagger:route PATCH /wallet/status/{Id} Wallet walletStatusUpdate
// Api used to update the status of a wallet, based on given wallet id.
// responses:
//   204: walletBalanceStatusResponse

// The status code is the only response.
// swagger:response walletBalanceStatusResponse
type walletBalanceStatusResponse struct{}

// swagger:parameters walletStatusUpdate
type walletStatusUpdateParamsWrapper struct {
	// Wallet id is passed as parameter.
	// in:body
	Body statusUpdateRequest
}

// swagger:route GET /wallet/status/{Id} Wallet walletStatusFetch
// Api used to fetch the status of a wallet, based on given wallet id.
// responses:
//   200: walletStatusFetchResponse

// The current status of the wallet is returned.
// swagger:response walletStatusFetchResponse
type walletStatusFetchResponse struct {
	// in:body
	Body statusFetchResponse
}

// swagger:parameters walletStatusFetch
type walletStatusFetchParamsWrapper struct {
	// Wallet id is passed as parameter.
	// in:path
	Id int
}
