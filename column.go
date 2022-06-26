package main

type Column struct {
	id                       int
	status                   string
	amountOfFloors           int
	amountOfElevators        int
	servedFloors             []int
	isBasement               bool
	callButtonList           []CallButton
	floor                    int
	direction                string
	elevatorList             []Elevator
	elevator                 Elevator
	bestElevatorInformations BestElevatorInformations
}

type BestElevatorInformations struct {
	bestScore    int
	referenceGap int
	bestElevator Elevator
}

//function to create the best elevator informations needed
func NewBestElevatorInformations(bestScore int, referenceGap int, bestElevator Elevator) BestElevatorInformations {

	bestElevatorInformations := BestElevatorInformations{bestScore: 6, referenceGap: 10000000, bestElevator: bestElevator}

	return bestElevatorInformations
}

//function to create a new column
func NewColumn(id int, status string, amountOfElevators int, servedFloors []int, _isBasement bool) *Column {

	column := &Column{id: id, status: "idle", amountOfElevators: amountOfElevators, servedFloors: servedFloors, isBasement: false}
	column.createCallButton(column.amountOfFloors, column.isBasement)
	column.createElevator(column.amountOfFloors, column.amountOfElevators)

	return column

}

//Simulate the creation of a button when a user is in the basement or on the other floors
func (c *Column) createCallButton(amountOfFloors int, isBasement bool) {
	if isBasement {
		var buttonFloor = -1
		for i := 0; i < amountOfFloors; i++ {
			if buttonFloor < amountOfFloors {
				callButton := NewCallButton(1, "OFF", buttonFloor, "up")
				c.callButtonList = append(c.callButtonList, callButton)
			}
			if buttonFloor > 1 {
				callButton := NewCallButton(i+1, "OFF", buttonFloor, "down")
				c.callButtonList = append(c.callButtonList, callButton)
			}
			buttonFloor++
		}
	}
}

//Simulate the creation of the elevator before it is requested
func (c *Column) createElevator(amountOfFloors int, amountOfElevators int) {

	for i := 0; i < amountOfElevators; i++ {
		elevator := NewElevator(i+1, "idle", amountOfFloors, 1)
		c.elevatorList = append(c.elevatorList, elevator)
	}
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(userPosition int, direction string) Elevator {

	elevator := c.findElevator(userPosition, direction)
	c.elevator.addNewRequest(userPosition)
	c.elevator.move()
	c.elevator.addNewRequest(1)
	c.elevator.move()
	return elevator
}

//Function to find an elevator upon the situation of the user
func (c *Column) findElevator(requestedFloor int, requestedDirection string) Elevator {

	bestElevatorInformations := NewBestElevatorInformations(6, 10000000, c.elevatorList[0])

	if requestedFloor == 1 {

		for _, elevator := range c.elevatorList {

			//The elevator is at the lobby and already has some requests. It is about to leave but has not yet departed.
			if 1 == elevator.currentFloor && elevator.status == "stopped" {

				bestElevatorInformations = c.checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor)

				//The elevator is at the lobby and has no requests
			} else if 1 == elevator.currentFloor && elevator.status == "idle" {

				bestElevatorInformations = c.checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)

				//The elevator is lower than me and is coming up. It means that I'm requesting an elevator to go to a basement, and the elevator is on it's way to me.
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {

				bestElevatorInformations = c.checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor)

				//The elevator is above me and is coming down. It means that I'm requesting an elevator to go to a floor, and the elevator is on it's way to me
			} else if 1 < elevator.currentFloor && elevator.direction == "down" {

				bestElevatorInformations = c.checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor)

				//The elevator is not at the first floor, but doesn't have any request
			} else if elevator.status == "idle" {

				bestElevatorInformations = c.checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor)

				//The elevator is not available, but still could take the call if nothing better is found

			} else {

				bestElevatorInformations = c.checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor)
			}
		}

	} else {

		for _, elevator := range c.elevatorList {
			//foreach (Elevator elevator in this.elevatorsList)
			//The elevator is at the same level as me, and is about to depart to the first floor
			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
				bestElevatorInformations = c.checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor)

				//The elevator is lower than me and is going up. I'm on a basement, and the elevator can pick me up on it's way
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == "up" {

				bestElevatorInformations = c.checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)

				//The elevator is higher than me and is going down. I'm on a floor, and the elevator can pick me up on it's way
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down" {

				bestElevatorInformations = c.checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)

				//The elevator is idle and has no requests
			} else if elevator.status == "idle" {

				bestElevatorInformations = c.checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor)
			} else {
				//The elevator is not available, but still could take the call if nothing better is found

				bestElevatorInformations = c.checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor)
			}
		}
	}
	return bestElevatorInformations.bestElevator
}

//Function to evaluate the best option of the available elevators upon ther user's request
func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator Elevator, bestElevatorInformations BestElevatorInformations, floor int) BestElevatorInformations {

	if scoreToCheck < bestElevatorInformations.bestScore {
		bestElevatorInformations.bestScore = scoreToCheck
		bestElevatorInformations.bestElevator = newElevator
		bestElevatorInformations.referenceGap = Abs(newElevator.currentFloor - floor)

	} else if bestElevatorInformations.bestScore == scoreToCheck {
		gap := Abs(newElevator.currentFloor - floor)
		if bestElevatorInformations.referenceGap > gap {

			bestElevatorInformations.bestElevator = newElevator
			bestElevatorInformations.referenceGap = gap
		}
	}
	return bestElevatorInformations
}
