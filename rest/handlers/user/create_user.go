package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommerce/domain"
	"ecommerce/utils"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := domain.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}
	createdUser, err := h.userRepo.Create(newUser)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	utils.SendData(w, createdUser, http.StatusCreated)
}
