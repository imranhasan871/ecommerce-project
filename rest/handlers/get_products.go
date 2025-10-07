package handlers

import (
	"net/http"

	"ecommerce/database"
	"ecommerce/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.List(), http.StatusOK)
}
