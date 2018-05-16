package model

//Dog holds all the information in regards to a users pet
type Dog struct {
	ID       string `json:"-"`
	Name     string `json:"name,omitempty"`
	Breed    string `json:"breed,omitempty"`
	Sex      string `json:"sex,omitempty"`
	ImageURL string `json:"image-url,omitempty"`
}
