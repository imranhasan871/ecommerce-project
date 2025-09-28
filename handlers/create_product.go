package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommerce/product"
	"ecommerce/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := product.Product{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid Json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, newProduct)

	w.WriteHeader(http.StatusCreated)

	utils.SendData(w, newProduct, http.StatusOK)
}
