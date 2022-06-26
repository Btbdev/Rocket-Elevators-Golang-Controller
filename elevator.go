package main

import "sort"

//Attributes of the elevator
type Elevator struct {
	id                    int
	status                string
	amountOfFloors        int
	currentFloor          int
	floorRequestsList     []int
	direction             string
	overweight            bool
	completedRequestsList []int
	screenDisplay         int
	door                  Door
}

//Function to create the elevator
func NewElevator(id int, status string, amountOfFloor int, currentFloor int) Elevator {

	//floorRequestList := floorRequestList{[]int}
	door := Door{id: id, status: "closed", obstruction: false}
	elevator := Elevator{id: id, status: "idle", amountOfFloors: amountOfFloor, currentFloor: currentFloor, door: door, floorRequestsList: []int{}, completedRequestsList: []int{}}

	return elevator

}

//Function to move the elevator upon the request made by the user
func (e *Elevator) move() {

	for true {
		if len(e.floorRequestsList) != 0 {
			var destination = e.floorRequestsList[0]
			e.status = "moving"
			if e.currentFloor < destination {
				e.direction = "up"
				e.sortFloorList()
				if e.currentFloor < destination {
					e.currentFloor++
					e.screenDisplay = e.currentFloor
				}
				break
			}
			if e.currentFloor > destination {
				e.direction = "down"
				e.sortFloorList()
				if e.currentFloor > destination {
					e.currentFloor--
					e.screenDisplay = e.currentFloor
				}
				break
			}
			e.status = "stopped"
			e.operateDoors()
			e.completedRequestsList = append(e.completedRequestsList, e.floorRequestsList[0])
			e.floorRequestsList = append([]int{1}, e.floorRequestsList...)
		}
		e.status = "idle"
	}
}

//Function to sort the floor list created
func (e *Elevator) sortFloorList() {

	if e.direction == "up" {
		sort.Ints(e.floorRequestsList)
	}
	if e.direction == "down" {
		sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
	}
}

//Function to operate the elevator's doors upon its situation
func (e *Elevator) operateDoors() {

	e.door.status = "opened"
	//Console.WriteLine("wait 5 seconds");
	if e.overweight == false {
		e.door.status = "closing"
		if e.door.obstruction == false {
			e.door.status = "closed"
		}
	}
	if e.door.obstruction == true {
		e.operateDoors()
	}
	if e.overweight == true {
		//Console.WriteLine("Activate overweight alarm")

		e.operateDoors()
	}
}

//Function to add the new requested floor on the list
func (e *Elevator) addNewRequest(requestedFloor int) {

	if len(e.floorRequestsList) == 0 {
		e.floorRequestsList = append(e.floorRequestsList, requestedFloor)
	}
	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > requestedFloor {
		e.direction = "down"
	}
}
