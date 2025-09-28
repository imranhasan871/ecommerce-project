package handlers

import (
	"net/http"

	"ecommerce/product"
	"ecommerce/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, product.ProductList, http.StatusOK)
}
