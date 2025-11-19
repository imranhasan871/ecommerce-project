package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommerce/domain"
	"ecommerce/utils"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	newProduct := domain.Product{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid Json", http.StatusBadRequest)
		return
	}
	createdProduct, err := h.productRepo.Create(newProduct)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}
	utils.SendData(w, createdProduct, http.StatusCreated)
}
