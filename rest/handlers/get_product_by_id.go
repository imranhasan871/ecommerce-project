package handlers

import (
	"net/http"
	"strconv"

	"ecommerce/database"
	"ecommerce/utils"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give a valid product id", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			utils.SendData(w, product, http.StatusOK)
			return
		}
	}
	utils.SendData(w, "Data Not Found", http.StatusNotFound)
}
