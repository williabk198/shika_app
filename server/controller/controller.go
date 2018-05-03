package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

type credentials struct {
	AllowedApps []struct {
		Name string
		Key  string
	} `json:"apps"`
	Salt string
}

var (
	//User holds the controller for the user
	User user
	//Dog hold the controller for the dog
	Dog dog

	creds credentials

	errAppDisallowed = errors.New("The application is not permitted to access this API")
	errBadKey        = errors.New("The appilication did not give the correct key")
)

func init() {
	User = user{}
	Dog = dog{}

	jsonFile, err := ioutil.ReadFile("../creds/api_keys.json")
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

	var appFound bool

	for _, appKey := range creds.AllowedApps {
		if appKey.Name == appName {
			appFound = true

			hasher := sha256.New()
			hasher.Write([]byte(key + creds.Salt))
			hash := hex.EncodeToString(hasher.Sum(nil))

			fmt.Printf("%s\n%s\n\n", hash, appKey.Key)

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
