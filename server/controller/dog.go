package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/williabk198/shika_app/server/model"

	"github.com/gorilla/mux"
	"github.com/williabk198/shika_app/server/db"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type dog struct{}

func (d dog) Get(w http.ResponseWriter, r *http.Request) {
	//TODO: Verify that Shika app is send the request and
	//write the information for the desired dog

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

	dogKey := mux.Vars(r)["dogKey"]
	dog, err := db.Dogs.Get(dogKey)
	if err != nil {
		log.Errorf(ctx, "Failed to retrieve dog with ID %s from the database\n%v",
			dogKey,
			err,
		)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(dog)
	if err != nil {
		log.Errorf(ctx, "Could not write GET Dog response")
	}
}

func (d dog) Post(w http.ResponseWriter, r *http.Request) {

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

	newDog := new(model.Dog)
	newDog.Name = r.FormValue("name")
	newDog.Breed = r.FormValue("breed")
	newDog.Sex = r.FormValue("sex")
	newDog.ImageURL = r.FormValue("image-url")

	dogKey, err := db.Dogs.Add(newDog)
	if err != nil {
		log.Errorf(ctx, "Unable to add dog to database\n%v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, dogKey)
}
