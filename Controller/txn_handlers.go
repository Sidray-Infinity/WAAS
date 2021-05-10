package Controller

import (
	"net/http"
	"waas/Domain"
)

func generateCSV(w http.ResponseWriter, r *http.Request) {
	Domain.GenerateCSV()
}
