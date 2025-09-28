package main

import (
	"fmt"
	"net/http"
	"os"

	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"ecommerce/product"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", (http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", (http.HandlerFunc(handlers.CreateProduct)))

	fmt.Println("Server is running on port 3000")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", middleware.Logger(globalRouter))
	if err != nil {
		fmt.Println("Error starting the server")
		os.Exit(1)
	}
}

func init() {
	product1 := product.Product{
		ID:          1,
		Title:       "Wireless Headphones",
		Description: "High-quality noise-cancelling headphones.",
		Price:       129.99,
		ImgURL:      "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS",
	}
	product2 := product.Product{
		ID:          2,
		Title:       "Smart Watch",
		Description: "Stylish smart watch with health tracking.",
		Price:       199.99,
		ImgURL:      "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	}
	product3 := product.Product{
		ID:          3,
		Title:       "Running Shoes",
		Description: "Lightweight shoes for everyday running.",
		Price:       89.50,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	}
	product4 := product.Product{
		ID:          4,
		Title:       "Wireless Headphones",
		Description: "High-quality noise-cancelling headphones.",
		Price:       129.99,
		ImgURL:      "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS",
	}
	product5 := product.Product{
		ID: 5, Title: "Smart Watch",
		Description: "Stylish smart watch with health tracking.",
		Price:       199.99,
		ImgURL:      "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	}
	product6 := product.Product{
		ID:          6,
		Title:       "Running Shoes",
		Description: "Lightweight shoes for everyday running.",
		Price:       89.50,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	}

	product.ProductList = append(product.ProductList, product1)
	product.ProductList = append(product.ProductList, product2)
	product.ProductList = append(product.ProductList, product3)
	product.ProductList = append(product.ProductList, product4)
	product.ProductList = append(product.ProductList, product5)
	product.ProductList = append(product.ProductList, product6)
}
