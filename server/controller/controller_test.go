package controller

import "testing"

func TestAPICheck(t *testing.T) {

	err := checkAPIKey("shika", "0sdfgjZLKJw459klsd345kjjLKJf956B")
	if err == nil {
		t.Error("Expected an error but did not get one")
	}
	if err != errBadKey {
		t.Errorf("Expected an errBadKey error but got another error\n%v", err)
	}

	err = checkAPIKey("bad_app", "M53BdwLOG0hlNJXK7cQpkSWOYMlC5UKW")
	if err == nil {
		t.Error("Expected an error but did not get one")
	}
	if err != errAppDisallowed {
		t.Errorf("Expected an errAppDisallowed error but got another error\n%v", err)
	}

	err = checkAPIKey("shika", "M53BdwLOG0hlNJXK7cQpkSWOYMlC5UKW")
	if err != nil {
		t.Errorf("Got an error but was not expecing one\n%v", err)
	}
}
