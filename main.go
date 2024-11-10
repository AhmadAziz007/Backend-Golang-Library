package main

import (
	"library-synapsis/app"
	"library-synapsis/controller"
	"library-synapsis/controller/implement3"
	"library-synapsis/helper"
	"library-synapsis/middleware"
	"library-synapsis/repository/implement"
	"library-synapsis/service/implement2"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.ConnectDB()
	validate := validator.New()

	//category
	categoryRepository := implement.NewCategoryRepository()
	categoryService := implement2.NewCategoryService(categoryRepository, db, validate)
	categoryController := implement3.NewCategoryController(categoryService)

	//author management
	authorManagementRepository := implement.NewAuthorManagementRepository()
	authorManagementService := implement2.NewAuthorManagementService(authorManagementRepository, db, validate)
	authorManagementController := implement3.NewAuthorManagementController(authorManagementService)

	//book management
	bookManagementRepository := implement.NewBookManagementRepository()
	bookManagementService := implement2.NewBookManagementService(bookManagementRepository, db, validate)
	bookManagementController := implement3.NewBookManagementController(bookManagementService)

	// Asumsi db adalah objek *sql.DB yang sudah terhubung ke database
	loginController := controller.NewLoginController(UserManagementRepository, db)

	router := app.NewRouter(categoryController, authorManagementController, bookManagementController, loginController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.AuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
