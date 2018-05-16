package model

//User holds the information that of the person visiting the dog park
type User struct {
	ID       string   `json:"-"`
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email,omitempty"`
	ImageURL string   `json:"image,omitempty"`
	Dogs     []string `json:"dogs,omitempty"`
}
