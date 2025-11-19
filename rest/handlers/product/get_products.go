package product

import (
	"net/http"

	"ecommerce/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to get products")
		return
	}
	utils.SendData(w, products, http.StatusOK)
}
