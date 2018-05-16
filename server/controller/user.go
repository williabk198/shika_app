package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williabk198/shika_app/server/db"
	"github.com/williabk198/shika_app/server/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type user struct{}

func (u user) Get(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	appKeyString := r.Header.Get("Authorization")
	if appKeyString == "" {
		log.Errorf(ctx, "API Key was not provided")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := checkAPIKey(getAppNameAndKey(appKeyString))
	if err != nil {
		log.Errorf(ctx, "API key authentication failed\n%v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userKey := mux.Vars(r)["userKey"]
	user, err := db.Dogs.Get(userKey)
	if err != nil {
		log.Errorf(ctx, "Failed to retrieve dog with ID %s from the database\n%v",
			userKey,
			err,
		)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Errorf(ctx, "Could not write GET Dog response")
	}
}

func (u user) Post(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	appKeyString := r.Header.Get("Authorization")
	if appKeyString == "" {
		log.Errorf(ctx, "API Key was not provided")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := checkAPIKey(getAppNameAndKey(appKeyString))
	if err != nil {
		log.Errorf(ctx, "API key authentication failed\n%v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newUser := new(model.User)
	newUser.Name = r.FormValue("name")
	newUser.Email = r.FormValue("email")
	newUser.ImageURL = r.FormValue("image-url")

	userKey, err := db.Users.Add(newUser)
	if err != nil {
		log.Errorf(ctx, "Unable to add dog to database\n%v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, userKey)
}
