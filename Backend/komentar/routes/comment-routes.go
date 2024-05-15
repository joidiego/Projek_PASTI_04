package routes

import (
	controllers "komentar/Controllers"

	"github.com/gorilla/mux"
)

var RegisterCommentRoutes = func(router *mux.Router) {
	router.HandleFunc("/comment/", controllers.CreateComment).Methods("POST")
	router.HandleFunc("/comment/", controllers.GetComment).Methods("GET")
	router.HandleFunc("/comment/{commentId}", controllers.GetCommentById).Methods("GET")
	router.HandleFunc("/comment/{commentId}", controllers.UpdateComment).Methods("PUT")
	router.HandleFunc("/comment/{commentId}", controllers.DeleteComment).Methods("DELETE")
}
