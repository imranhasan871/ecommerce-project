package main

import "net/http"

func getProducts(w http.ResponseWriter, r *http.Request) {
	sendData(w, ProductList, http.StatusOK)
}
