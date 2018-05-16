package model

//Visitor holds which user and which dog(s)
//the user has checked in along with the area
//that the user checked in to.
type Visitor struct {
	ID          string   `json:"-"`
	UserKey     string   `json:"user"`
	DogKeys     []string `json:"dogs"`
	CheckInTime string   `json:"check-in"`
	ParkArea    string   `json:"park-area"`
}
