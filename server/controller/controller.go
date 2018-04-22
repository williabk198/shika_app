package controller

var (
	//User holds the controller for the user
	User user
	//Dog hold the controller for the dog
	Dog dog
)

func init() {
	User = user{}
	Dog = dog{}
}
