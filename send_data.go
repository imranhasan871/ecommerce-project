package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendData(w http.ResponseWriter, data any, statusCode int) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid Json", statusCode)
		return
	}
}
