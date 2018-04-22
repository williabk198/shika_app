package controller

import "net/http"

type dog struct{}

func (d dog) Get(w http.ResponseWriter, r *http.Request) {
	//TODO: Verify that Shika app is send the request and
	//write the information for the desired dog
	w.WriteHeader(http.StatusNotImplemented)
}

func (d dog) Post(w http.ResponseWriter, r *http.Request) {
	//TODO: Verify that Shika app is send the request and
	//store the information recieved from the app
	w.WriteHeader(http.StatusNotImplemented)
}
