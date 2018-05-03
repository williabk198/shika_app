package controller

import "net/http"

type user struct{}

func (u user) VerifiedGet(w http.ResponseWriter, r *http.Request) {
	//TODO: Verify that Shika app is send the request and
	//write the information for the desired user
	w.WriteHeader(http.StatusNotImplemented)
}

func (u user) Post(w http.ResponseWriter, r *http.Request) {
	//TODO: Verify that Shika app is send the request and
	//store the information recieved from the app
	w.WriteHeader(http.StatusNotImplemented)
}
