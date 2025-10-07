package handlers

import (
	"net/http"
	"strconv"

	"ecommerce/database"
	"ecommerce/utils"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give a valid product id", http.StatusBadRequest)
		return
	}

	product := database.Get(id)
	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product Not Found")
		return
	}

	utils.SendData(w, product, http.StatusOK)
}
