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
	elevator          Elevator
	elevatorList      []Elevator

	//createElevator(amountOfFloors amountOfElevator)
	//createCallButtons(amountOfFloors isBasement)

}

func NewColumn(id int, status string, amountOfElevators int, servedFloors []int, _isBasement bool) *Column {

	//callButton := CallButton{id: id, status: status, floor: floor, direction: direction}
	//elevator := Elevator{id: id, status: status, amountOfFloors: amountOfFloors, currentFloor: currentFloor, door: door}
	column := &Column{id: id, status: "idle", amountOfElevators: amountOfElevators, servedFloors: servedFloors, isBasement: false}

	return column
}

// func (c *Column) createCallButton(amountOfFloors int, isBasement bool) *Elevator {
// 	if isBasement {
// 		var buttonFloor = -1
// 		for i := 0; i < amountOfFloors; i++ {
// 			if buttonFloor < amountOfFloors {
// 				var callButton = NewCallButton(i+1, "OFF", 1, "up")
// 				c.callButtonList = append(c.callButtonList, callButton)
// 			}
// 			if buttonFloor > 1 {
// 				var callButton = NewCallButton(i+1, "OFF", buttonFloor, "down")
// 				c.callButtonList = append(c.callButtonList, callButton)
// 			}
// 			buttonFloor++
// 		}
// 	}
// }

// func (c *Column) createElevator(amountOfFloors int, amountOfElevators int) {

// 	for i := 0; i < amountOfElevators; i++ {
// 		var elevator = NewElevator(i+1, "idle", amountOfFloors, 1)
// 		c.elevatorList = append(c.elevatorList, elevator)
// 	}
// }

// //Simulate when a user press a button on a floor to go back to the first floor
// func (c *Column) requestElevator(userPosition int, direction string) *Elevator {

// 	var elevator = findElevator(userPosition, direction)
// 	elevator.addNewRequest(userPosition)
// 	elevator.move()
// 	elevator.addNewRequest(1)
// 	elevator.move()
// 	return elevator
// }

func (c *Column) findElevator(requestedFloor int, requestedDirection string) {

			var bestElevatorInformations = nil
            //Elevator bestElevator = null;

            if (requestedFloor == 1) {

                //for i, s := range Elevator c.elevatorsList {

                    //The elevator is at the lobby and already has some requests. It is about to leave but has not yet departed.
                    if 1 == elevator.currentFloor && elevator.status == "stopped" {

                        bestElevatorInformations = checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor);

                    }
                    //The elevator is at the lobby and has no requests
                    if 1 == elevator.currentFloor && elevator.status == "idle" {
                        
						bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor);

                    }
                    //The elevator is lower than me and is coming up. It means that I'm requesting an elevator to go to a basement, and the elevator is on it's way to me.
                    if 1 > elevator.currentFloor && elevator.direction == "up" {
                        
						bestElevatorInformations = checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor);

                    }
                    //The elevator is above me and is coming down. It means that I'm requesting an elevator to go to a floor, and the elevator is on it's way to me
                    if 1 < elevator.currentFloor && elevator.direction == "down" {
                        
						bestElevatorInformations = checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor);

                    }
                    //The elevator is not at the first floor, but doesn't have any request
                    if elevator.status == "idle" {
                        
						bestElevatorInformations = checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor);

                    }
                    //The elevator is not available, but still could take the call if nothing better is found
                    
					else 
                        
						bestElevatorInformations = checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor);
                    
                };
            }
            else {

                foreach (Elevator elevator in this.elevatorsList)
                {
                    //The elevator is at the same level as me, and is about to depart to the first floor
                    if (requestedFloor == elevator.currentFloor && elevator.status == "stopped" && _requestedDirection == elevator.direction)
                    {
                        bestElevatorInformations = checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor);
                    }
                    //The elevator is lower than me and is going up. I'm on a basement, and the elevator can pick me up on it's way
                    else if (requestedFloor > elevator.currentFloor && elevator.direction == "up" && _requestedDirection == "up")
                    {
                        bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor);
                    }
                    //The elevator is higher than me and is going down. I'm on a floor, and the elevator can pick me up on it's way
                    else if (requestedFloor < elevator.currentFloor && elevator.direction == "down" && _requestedDirection == "down")
                    {
                        bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor);
                    }
                    //The elevator is idle and has no requests
                    else if (elevator.status == "idle")
                    {
                        bestElevatorInformations = checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor);
                    }
                    //The elevator is not available, but still could take the call if nothing better is found
                    else
                    {
                        bestElevatorInformations = checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor);
                    }
                };
            }
                bestElevator = bestElevatorInformations.bestElevator;
                return bestElevator;
}
