package cmd

import (
	"fmt"
	"net/http"
	"os"

	"ecommerce/middleware"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.CorsWithPreflight)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", manager.With(mux))
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
