package app

import (
	"context"
	"library-synapsis/controller"
	"library-synapsis/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandlerToHTTPRouter(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Tambahkan params ke dalam context, jika perlu
		//ctx := r.Context()
		//ctx = context.WithValue(ctx, "params", ps)
		//r = r.WithContext(ctx)
		//h.ServeHTTP(w, r)
		ctx := context.WithValue(r.Context(), "params", ps)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

func WrapHandlerWithParams(handle httprouter.Handle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Mengambil params dari httprouter dan meneruskannya ke handler
		ps := httprouter.ParamsFromContext(r.Context())
		handle(w, r, ps)
	}
}

func NewRouter(categoryController controller.CategoryController,
	authorManagementController controller.AuthorManagementController,
	bookManagementController controller.BookManagementController,
	userManagementController controller.UserManagementController,
	stockManagementController controller.StockManagementController,
	rentBookController controller.RentBookController,
	//loginController controller.LoginController,
) *httprouter.Router {
	router := httprouter.New()

	//category
	router.GET("/api/categories", categoryController.FindAllCategory)
	router.GET("/api/categories/:categoryId", categoryController.FindByCategoryId)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	//User Management
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
	//router.POST("/api/login", loginController.Login)

	//Book Management
	//router.GET("/api/books", HandlerToHTTPRouter(middleware.AuthMiddleware(WrapHandlerWithParams(bookManagementController.FindAllBook), "Admin1", "Admin2")))
	//router.POST("/api/books", HandlerToHTTPRouter(middleware.AuthMiddleware(WrapHandlerWithParams(bookManagementController.FindByBookLikeCriteria), "Admin1", "Admin2", "Customer")))
	//router.GET("/api/books/:bookId", HandlerToHTTPRouter(middleware.AuthMiddleware(WrapHandlerWithParams(bookManagementController.FindByBookId), "Admin1", "Admin2")))
	//router.POST("/api/books/create", HandlerToHTTPRouter(middleware.AuthMiddleware(WrapHandlerWithParams(bookManagementController.CreateBook), "Admin1", "Admin2")))
	//router.PUT("/api/books/update/:bookId", HandlerToHTTPRouter(middleware.AuthMiddleware(WrapHandlerWithParams(bookManagementController.UpdateBook), "Admin1", "Admin2")))
	//router.DELETE("/api/books/delete/:bookId", HandlerToHTTPRouter(middleware.AuthMiddleware(WrapHandlerWithParams(bookManagementController.DeleteBook), "Admin1", "Admin2")))

	router.GET("/api/books", bookManagementController.FindAllBook)
	router.POST("/api/books", bookManagementController.FindByBookLikeCriteria)
	router.GET("/api/books/:bookId", bookManagementController.FindByBookId)
	router.POST("/api/books/create", bookManagementController.CreateBook)
	router.PUT("/api/books/update/:bookId", bookManagementController.UpdateBook)
	router.DELETE("/api/books/delete/:bookId", bookManagementController.DeleteBook)

	//Stock Management
	router.GET("/api/stocks", stockManagementController.FindAllStock)
	router.GET("/api/stocks/:stockId", stockManagementController.FindByStockId)
	router.POST("/api/stocks/create", stockManagementController.CreateStock)
	router.PUT("/api/stocks/update/:stockId", stockManagementController.UpdateStock)
	router.DELETE("/api/stocks/delete/:stockId", stockManagementController.DeleteStock)

	//Rent Book
	router.GET("/api/rents", rentBookController.FindAllRent)
	router.GET("/api/rents/:rentId", rentBookController.FindByRentId)
	router.POST("/api/rents/create", rentBookController.CreateRent)
	router.PUT("/api/rents/update/:rentId", rentBookController.UpdateRent)
	router.DELETE("/api/rents/delete/:rentId", rentBookController.DeleteRent)

	router.PanicHandler = exception.ErrorHandler

	return router
}
