package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	id        int
	status    string
	floor     int
	direction string
}

func NewFloorRequestButton(id int, status string, floor int, direction string) *FloorRequestButton {

	floorRequestButton := &FloorRequestButton{id: id, status: "idle", floor: floor, direction: "up"}

	return floorRequestButton
}
