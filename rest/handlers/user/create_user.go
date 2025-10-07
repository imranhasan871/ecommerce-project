package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommerce/database"
	"ecommerce/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := database.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}
	createdUser := newUser.Store()
	utils.SendData(w, createdUser, http.StatusCreated)
}
