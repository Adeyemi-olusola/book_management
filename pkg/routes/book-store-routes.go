package routes

import (
	//"book-management/pkg/controllers/book-controller"
	"book_management/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoresRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/book/{bookId}" , controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/book/{bookid}", controller.DeleteBook).Methods("DELETE")

}
