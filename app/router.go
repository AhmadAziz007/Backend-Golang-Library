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
	BookManagementController controller.BookManagementController, loginController controller.LoginController) *httprouter.Router {
	router := httprouter.New()

	//category
	router.GET("/api/categories", categoryController.FindAllCategory)
	router.GET("/api/categories/:categoryId", categoryController.FindByCategoryId)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

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
	router.GET("/api/books", middleware.AuthMiddleware(http.HandlerFunc(BookManagementController.FindAllBook), "Admin1", "Admin2", "Customer"))
	router.POST("/api/books", middleware.AuthMiddleware(http.HandlerFunc(BookManagementController.FindByBookLikeCriteria), "Admin1", "Admin2", "Customer"))
	router.GET("/api/books/:bookId", middleware.AuthMiddleware(http.HandlerFunc(BookManagementController.FindByBookId), "Admin1", "Admin2"))
	router.POST("/api/books/create", middleware.AuthMiddleware(http.HandlerFunc(BookManagementController.CreateBook), "Admin1", "Admin2"))
	router.PUT("/api/books/update/:bookId", middleware.AuthMiddleware(http.HandlerFunc(BookManagementController.UpdateBook), "Admin1", "Admin2"))
	router.DELETE("/api/books/delete/:bookId", middleware.AuthMiddleware(http.HandlerFunc(BookManagementController.DeleteBook), "Admin1", "Admin2"))

	router.PanicHandler = exception.ErrorHandler

	return router
}
