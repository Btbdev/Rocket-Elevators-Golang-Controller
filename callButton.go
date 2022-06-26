package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	id        int
	status    string
	floor     int
	direction string
}

//Function to create the callButton
func NewCallButton(id int, status string, floor int, direction string) CallButton {

	callButton := CallButton{id: id, status: status, floor: floor, direction: direction}

	return callButton
}
