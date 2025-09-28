package cmd

import (
	"fmt"
	"net/http"
	"os"

	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", (http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", (http.HandlerFunc(handlers.CreateProduct)))

	fmt.Println("Server is running on port 3000")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", middleware.Logger(globalRouter))
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
