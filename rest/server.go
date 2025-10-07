package rest

import (
	"fmt"
	"net/http"
	"os"

	"ecommerce/config"
	"ecommerce/rest/middlewares"
)

func Start() {
	cnf := config.GetConfig()

	manager := middlewares.NewManager()
	manager.Use(middlewares.Logger, middlewares.CorsWithPreflight)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	fmt.Println("Server is running on port:", cnf.HttpPort)
	err := http.ListenAndServe(":"+cnf.HttpPort, manager.With(mux))
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
