package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommerce/database"
	"ecommerce/utils"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}
	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	utils.SendData(w, "Login Successfully", http.StatusCreated)
}
