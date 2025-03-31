package routes

import (
	controllers "github.com/garvit4540/go-postgres-bookstore/pkg/controllers"
	"github.com/garvit4540/go-postgres-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods(utils.MethodPOST)
	router.HandleFunc("/book/", controllers.GetBook).Methods(utils.MethodGET)
	router.HandleFunc("/book/{book_id}", controllers.GetBookById).Methods(utils.MethodGET)
	router.HandleFunc("/book/{book_id}", controllers.UpdateBook).Methods(utils.MethodPUT)
	router.HandleFunc("/book/{book_id}", controllers.DeleteBook).Methods(utils.MethodDELETE)
}
