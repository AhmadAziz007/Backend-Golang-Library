package app

import (
	"library-synapsis/controller"
	"library-synapsis/exception"
	"library-synapsis/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController,
	authorManagementController controller.AuthorManagementController,
	bookManagementController controller.BookManagementController,
	userManagementController controller.UserManagementController,
	loginController controller.LoginController) *httprouter.Router {
	router := httprouter.New()

	//category
	router.GET("/api/categories", categoryController.FindAllCategory)
	router.GET("/api/categories/:categoryId", categoryController.FindByCategoryId)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/users", userManagementController.FindAllUser)
	router.GET("/api/users/:userId", userManagementController.FindByUserId)
	router.POST("/api/users/create", userManagementController.CreateUser)
	router.PUT("/api/users/update/:userId", userManagementController.UpdateUser)
	router.DELETE("/api/users/delete/:userId", userManagementController.DeleteUser)

	//Author Management
	router.GET("/api/authors", authorManagementController.FindAllAuthor)
	router.GET("/api/authors/:authorId", authorManagementController.FindByAuthorId)
	router.POST("/api/authors", authorManagementController.CreateAuthor)
	router.PUT("/api/authors/:authorId", authorManagementController.UpdateAuthor)
	router.DELETE("/api/authors/:authorId", authorManagementController.DeleteAuthor)

	// Login route
	router.POST("/api/login", loginController.Login)

	//Book Management
	// Membungkus controller methods menjadi http.Handler
	router.GET("/api/books", middleware.AuthMiddleware(http.HandlerFunc(bookManagementController.FindAllBook), "Admin1", "Admin2", "Customer"))
	router.POST("/api/books", middleware.AuthMiddleware(http.HandlerFunc(bookManagementController.FindByBookLikeCriteria), "Admin1", "Admin2", "Customer"))
	router.GET("/api/books/:bookId", middleware.AuthMiddleware(http.HandlerFunc(bookManagementController.FindByBookId), "Admin1", "Admin2"))
	router.POST("/api/books/create", middleware.AuthMiddleware(http.HandlerFunc(bookManagementController.CreateBook), "Admin1", "Admin2"))
	router.PUT("/api/books/update/:bookId", middleware.AuthMiddleware(http.HandlerFunc(bookManagementController.UpdateBook), "Admin1", "Admin2"))
	router.DELETE("/api/books/delete/:bookId", middleware.AuthMiddleware(http.HandlerFunc(bookManagementController.DeleteBook), "Admin1", "Admin2"))

	router.PanicHandler = exception.ErrorHandler

	return router
}
