package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// "ecommerce/cmd"

func main() {
	// cmd.Serve()

	// var s string
	// s = "Imran Hasan"

	// byteArr := []byte(s)
	// fmt.Println(s)
	// fmt.Println(byteArr)

	// base64.URLEncoding.WithPadding(base64.NoPadding)

	// data := []byte("Hello World")

	// hash := sha256.Sum256(data)
	// fmt.Printf("hash: %v\n", hash)

	secret := []byte("My-Secret")
	message := []byte("Hello World")

	h := hmac.New(sha256.New, secret)
	h.Write(message)

	text := h.Sum(nil)
	fmt.Println(text)

}
