package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommerce/database"
	"ecommerce/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := database.Product{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid Json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	w.WriteHeader(http.StatusCreated)

	utils.SendData(w, newProduct, http.StatusOK)
}
