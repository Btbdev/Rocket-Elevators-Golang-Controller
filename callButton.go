package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	id        int
	status    string
	floor     int
	direction string
}

func NewCallButton(id int, status string, floor int, direction string) *CallButton {

	CallButton := &CallButton{id: id, status: "idle", floor: floor, direction: "down"}

	return CallButton
}
