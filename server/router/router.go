package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williabk198/shika_app/server/controller"
)

func init() {

	router := mux.NewRouter()

	router.HandleFunc("/user/{userKey}", controller.User.VerifiedGet).Methods("POST")
	router.HandleFunc("/user", controller.User.Post).Methods("POST")

	router.HandleFunc("/dog/{dogKey}", controller.Dog.VerifiedGet).Methods("POST")
	router.HandleFunc("/dog", controller.Dog.Post).Methods("POST")

	http.Handle("/", router)
}
