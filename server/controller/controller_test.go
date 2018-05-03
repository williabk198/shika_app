package controller

import "testing"

func TestAPICheck(t *testing.T) {

	testApp := struct {
		Name string
		Key  string
	}{
		"test",
		"b6d637a20cc423665dc49b5c2bd008c28c1f2431d4f45f6ab4241ca2df5bec29",
	}

	creds.AllowedApps = append(creds.AllowedApps, testApp)
	creds.Salt = "76543210"

	err := checkAPIKey("test", "01234567")
	if err == nil {
		t.Error("Expected an error but did not get one")
	}
	if err != errBadKey {
		t.Errorf("Expected an errBadKey error but got another error\n%v", err)
	}

	err = checkAPIKey("bad_app", "12345678")
	if err == nil {
		t.Error("Expected an error but did not get one")
	}
	if err != errAppDisallowed {
		t.Errorf("Expected an errAppDisallowed error but got another error\n%v", err)
	}

	err = checkAPIKey("test", "12345678")
	if err != nil {
		t.Errorf("Got an error but was not expecing one\n%v", err)
	}
}
