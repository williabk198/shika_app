package router

import (
	"shika_app/server/controller"

	"github.com/gorilla/mux"
)

func init() {

	router := mux.NewRouter()

	router.HandleFunc("/user/{userKey}", controller.User.Get).Methods("GET")
	router.HandleFunc("/user", controller.User.Post).Methods("POST")

	router.HandleFunc("/dog/{dogKey}", controller.Dog.Get).Methods("GET")
	router.HandleFunc("/dog", controller.Dog.Post).Methods("POST")
}
