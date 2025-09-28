package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := Product{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid Json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(ProductList) + 1
	ProductList = append(ProductList, newProduct)

	w.WriteHeader(http.StatusCreated)

	sendData(w, newProduct, http.StatusOK)
}
