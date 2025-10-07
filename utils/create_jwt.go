package utils

import "encoding/json"

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         string `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrayHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	byteArrayData, err := json.Marshal()
}
