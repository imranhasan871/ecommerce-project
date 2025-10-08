package product

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
	createdProduct := database.Store(newProduct)
	utils.SendData(w, createdProduct, http.StatusCreated)
}
