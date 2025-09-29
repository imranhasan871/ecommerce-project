package cmd

import (
	"fmt"
	"net/http"
	"os"

	"ecommerce/global_router"
	"ecommerce/middleware"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
