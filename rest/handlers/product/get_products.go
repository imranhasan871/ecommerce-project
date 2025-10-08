package product

import (
	"net/http"

	"ecommerce/database"
	"ecommerce/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.List(), http.StatusOK)
}
