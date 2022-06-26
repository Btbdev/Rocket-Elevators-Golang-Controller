package main

//door is the constructor of the program that create the doors for the elevator
type Door struct {
	id          int
	status      string
	obstruction bool
}

//Function to create the door
func NewDoor(id int, status string, obsturction bool) *Door {

	door := &Door{id: id, status: "closed", obstruction: false}

	return door
}
