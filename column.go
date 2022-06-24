package main

type Column struct {
	id                int
	status            string
	amountOfFloors    int
	amountOfElevators int
	servedFloors      []int
	isBasement        bool
	callButtonList    []int
	callButton        CallButton
	floor             int
	direction         string
	elevator          int

	//createElevator(amountOfFloor, amountOfElevator) *Column {}
	//createCallButtons(amountOfFloor, isBasement) * Column {}

}

func NewColumn(id int, status string, amountOfElevators int, servedFloors []int, _isBasement bool) *Column {

	callButton := CallButton{id: id, status: status, floor: floor, direction: direction}
	elevator := Elevator{id: id, status: status, amountOfFloors: amountOfFloor, currentFloor: currentFloor, door: door}
	column := &Column{id: id, status: "idle", amountOfElevators: amountOfElevators, servedFloors: servedFloors, isBasement: false, callButton: callButton}

	return column
}

// func (c *Column) createCallButton(amountOfFloors int, isBasement bool) {
// 	if isBasement {
// 		var buttonFloor = -1
// 		for i := 0; i < amountOfFloor; i++ {
// 			if buttonFloor < amountOfFloor {
// 				var CallButton = CallButton(i+1, "OFF", 1, "up")
// 				column.callButtonList = append(c.callButtonList, CallButton)
// 			}
// 			if buttonFloor > 1 {
// 				var CallButton CallButton = CallButton(i+1, "OFF", buttonFloor, "down")
// 				column.callButtonList = append(c.callButtonList, CallButton)
// 			}
// 			buttonFloor++
// 		}
// 	}
// }

func (c *Column) createElevator(amountOfFloor, amountOfElevator) {

	for i := 0; i < amountOfElevators; i++ {
		var elevator = Elevator(i+1, "idle", _amountOfFloor, 1)
		this.elevatorsList.Add(elevator)
	}
}

// //Simulate when a user press a button on a floor to go back to the first floor
// func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {

// }
