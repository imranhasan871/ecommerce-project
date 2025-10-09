package rest

import (
	"fmt"
	"net/http"
	"os"

	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start(cnf config.Config) {

	manager := middlewares.NewManager()
	manager.Use(middlewares.CorsWithPreflight, middlewares.Logger)

	mux := http.NewServeMux()

	// initRoutes(mux, manager)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	fmt.Println("Server is running on port:", server.cnf.HttpPort)
	err := http.ListenAndServe(":"+server.cnf.HttpPort, manager.With(mux))
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
