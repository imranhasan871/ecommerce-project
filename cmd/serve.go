package cmd

import (
	"fmt"
	"net/http"
	"os"

	"ecommerce/config"
	"ecommerce/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.CorsWithPreflight)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	fmt.Println("Server is running on port:", cnf.HttpPort)
	err := http.ListenAndServe(":"+cnf.HttpPort, manager.With(mux))
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
