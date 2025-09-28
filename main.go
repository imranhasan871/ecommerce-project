package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageURL"`
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

func handlePreflightReq(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func sendData(w http.ResponseWriter, data any, statusCode int) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid Json", statusCode)
		return
	}
}

var ProductList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	if r.Method == http.MethodOptions {
		handlePreflightReq(w)
	}

	sendData(w, ProductList, http.StatusOK)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	if r.Method == http.MethodOptions {
		handlePreflightReq(w)
	}

	newProduct := Product{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid Json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(ProductList) + 1
	ProductList = append(ProductList, newProduct)

	w.WriteHeader(http.StatusCreated)

	sendData(w, newProduct, http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", Logger(http.HandlerFunc(getProducts)))
	mux.Handle("POST /products", http.HandlerFunc(createProduct))
	mux.Handle("OPTIONS /products", http.HandlerFunc(createProduct))

	fmt.Println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server")
		os.Exit(1)
	}
}

func init() {
	product1 := Product{
		ID:          1,
		Title:       "Wireless Headphones",
		Description: "High-quality noise-cancelling headphones.",
		Price:       129.99,
		ImgURL:      "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS",
	}
	product2 := Product{
		ID:          2,
		Title:       "Smart Watch",
		Description: "Stylish smart watch with health tracking.",
		Price:       199.99,
		ImgURL:      "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	}
	product3 := Product{
		ID:          3,
		Title:       "Running Shoes",
		Description: "Lightweight shoes for everyday running.",
		Price:       89.50,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	}
	product4 := Product{
		ID:          4,
		Title:       "Wireless Headphones",
		Description: "High-quality noise-cancelling headphones.",
		Price:       129.99,
		ImgURL:      "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS",
	}
	product5 := Product{
		ID: 5, Title: "Smart Watch",
		Description: "Stylish smart watch with health tracking.",
		Price:       199.99,
		ImgURL:      "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	}
	product6 := Product{
		ID:          6,
		Title:       "Running Shoes",
		Description: "Lightweight shoes for everyday running.",
		Price:       89.50,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	}

	ProductList = append(ProductList, product1)
	ProductList = append(ProductList, product2)
	ProductList = append(ProductList, product3)
	ProductList = append(ProductList, product4)
	ProductList = append(ProductList, product5)
	ProductList = append(ProductList, product6)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s | %s", r.Method, r.URL.Path, time.Since(start))
	})
}
