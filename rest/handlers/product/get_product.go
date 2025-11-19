package product

import (
	"net/http"
	"strconv"

	"ecommerce/utils"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give a valid product id", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.Get(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to get product")
		return
	}
	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product Not Found")
		return
	}

	utils.SendData(w, product, http.StatusOK)
}
