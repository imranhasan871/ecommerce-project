package rest

import (
	"net/http"

	"ecommerce/rest/handlers"
	"ecommerce/rest/middlewares"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductByID)))
}
