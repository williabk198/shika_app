package model

//Event holds all the necissary information for a particular event
type Event struct {
	ID        string   `json:"-"`
	Name      string   `json:"name,omitempty"`
	Date      string   `json:"date,omitempty"`
	Length    string   `json:"length,omitempty"`
	TimeStart []string `json:"time-start,omitempty"`
	TimeEnd   []string `json:"time-end,omitempty"`
}
