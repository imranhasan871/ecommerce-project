package product

import (
	"net/http"
	"strconv"

	"ecommerce/utils"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Please give a valid product id", http.StatusBadRequest)
		return
	}

	err = h.productRepo.Delete(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	utils.SendData(w, "Successfully delete product", http.StatusCreated)
}
