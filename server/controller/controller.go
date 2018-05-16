package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

type credentials struct {
	AllowedApps []struct {
		Name string
		Key  string
	} `json:"apps"`
	Salt string
}

var (
	//User holds the controller for the user API endpoint
	User = user{}
	//Dog holds the controller for the dog API endpoint
	Dog = dog{}
	//Event holds the controller for the event API endpoint
	Event = event{}
	//Visitor holds rhw controller fort the visitor API endpoint
	Visitor = visitor{}

	creds credentials

	errAppDisallowed error
	errBadKey        error
)

func init() {

	jsonFile, err := ioutil.ReadFile("./creds/api_keys.json")
	if err != nil {
		log.Printf(
			"Could not access API Key information\n%v",
			err,
		)
	}

	err = json.Unmarshal(jsonFile, &creds)
	if err != nil {
		log.Printf(
			"Could not load API Key information\n%v",
			err,
		)
	}
}

func checkAPIKey(appName, key string) error {

	errAppDisallowed = errors.New("The \"" + appName + "\" application is not permitted to access this API")
	errBadKey = errors.New("The \"" + appName + "\" appilication did not give the correct key")

	var appFound bool

	for _, appKey := range creds.AllowedApps {
		if appKey.Name == appName {
			appFound = true

			hasher := sha256.New()
			hasher.Write([]byte(key + creds.Salt))
			hash := hex.EncodeToString(hasher.Sum(nil))

			if hash == appKey.Key {
				return nil
			}
		}
	}

	if appFound {
		return errBadKey
	}

	return errAppDisallowed
}

func getAppNameAndKey(authHeader string) (string, string) {
	spaceIndex := strings.Index(authHeader, " ")
	appName := authHeader[0:spaceIndex]
	apiKey := authHeader[spaceIndex+1 : len(authHeader)]

	return appName, apiKey
}
