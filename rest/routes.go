package rest

import (
	"net/http"

	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	// Users Routes
	mux.Handle("POST /users", manager.With(http.HandlerFunc(user.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(user.Login)))

	// Products Routes
	mux.Handle("GET /products", manager.With(http.HandlerFunc(product.GetProducts)))
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(product.CreateProduct),
			middlewares.AuthenticateJWT,
		),
	)

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(product.GetProduct),
		))

	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(
				product.UpdateProduct),
			middlewares.AuthenticateJWT,
		),
	)

	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(product.DeleteProduct),
			middlewares.AuthenticateJWT),
	)
}
