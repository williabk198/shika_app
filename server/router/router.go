package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williabk198/shika_app/server/controller"
)

func init() {

	router := mux.NewRouter()

	router.HandleFunc("/user/{userKey}", controller.User.Get).Methods("GET")
	router.HandleFunc("/user", controller.User.Post).Methods("POST")

	router.HandleFunc("/dog/{dogKey}", controller.Dog.Get).Methods("GET")
	router.HandleFunc("/dog", controller.Dog.Post).Methods("POST")

	router.HandleFunc("/event/{dogKey}", controller.Dog.Get).Methods("GET")
	router.HandleFunc("/event", controller.Dog.Post).Methods("POST")

	router.HandleFunc("/visitor/{dogKey}", controller.Dog.Get).Methods("GET")
	router.HandleFunc("/visitor", controller.Dog.Post).Methods("POST")

	http.Handle("/", router)
}
