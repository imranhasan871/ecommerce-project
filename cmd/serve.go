package cmd

import (
	"fmt"
	"os"

	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println("Failed to connect to database", err)
		os.Exit(1)
	}

	middlewares := middlewares.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(dbCon)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start(*cnf)
}
