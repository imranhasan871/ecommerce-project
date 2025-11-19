package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommerce/config"
	"ecommerce/utils"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	cnf := config.GetConfig()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	usr, err := h.userRepo.Find(reqLogin.Email, reqLogin.Password)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	accessToken, err := utils.CreateJWT(
		cnf.JwtSecretKey,
		utils.Payload{
			Sub:       usr.ID,
			FirstName: usr.FirstName,
			LastName:  usr.LastName,
		})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, accessToken, http.StatusCreated)
}
