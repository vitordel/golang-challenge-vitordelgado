package routes

import (
	"golang-challenge-vitordelgado/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	r.HandleFunc("/login", controllers.Login).Methods("POST")
}
