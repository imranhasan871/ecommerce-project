package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"ecommerce/domain"
	"ecommerce/utils"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Please give a valid product id", http.StatusBadRequest)
		return
	}

	updateProduct := domain.Product{}

	error := json.NewDecoder(r.Body).Decode(&updateProduct)
	if error != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid Json", http.StatusBadRequest)
		return
	}

	updateProduct.ID = id
	updatedProduct, err := h.productRepo.Update(updateProduct)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	utils.SendData(w, updatedProduct, http.StatusCreated)
}
