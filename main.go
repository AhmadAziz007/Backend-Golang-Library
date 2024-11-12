package main

import (
	"library-synapsis/app"
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

	userManagementRepository := implement.NewUserManagementRepository()
	userManagementService := implement2.NewUserManagementService(userManagementRepository, db, validate)
	userManagementController := implement3.NewUserManegementController(userManagementService)

	//book management
	bookManagementRepository := implement.NewBookManagementRepository()
	bookManagementService := implement2.NewBookManagementService(bookManagementRepository, db, validate)
	bookManagementController := implement3.NewBookManagementController(bookManagementService)

	//stock management
	stockManagementRepository := implement.NewStockManagementRepository()
	stockManagementService := implement2.NewStockManagementService(stockManagementRepository, db, validate)
	stockManagementController := implement3.NewStockManagementController(stockManagementService)

	//rent book
	rentBookRepository := implement.NewRentBookRepository()
	rentBookService := implement2.NewRentBookService(rentBookRepository, db, validate)
	rentBookController := implement3.NewRentBookController(rentBookService)

	// Asumsi db adalah objek *sql.DB yang sudah terhubung ke database
	//loginRepository := implement.NewLoginRepository()
	//loginService := implement2.NewLoginService(loginRepository, db, validate)
	//loginController := implement3.NewLoginController(loginService)

	router := app.NewRouter(
		categoryController,
		authorManagementController,
		bookManagementController,
		userManagementController,
		stockManagementController,
		rentBookController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
