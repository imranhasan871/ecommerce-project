package handlers

import (
	"net/http"
	"strconv"

	"ecommerce/database"
	"ecommerce/utils"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Please give a valid product id", http.StatusBadRequest)
		return
	}

	database.Delete(id)

	utils.SendData(w, "Successfully delete product", http.StatusCreated)
}
