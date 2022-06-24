package main

type Elevator struct {
	id                    int
	status                string
	amountOfFloor         int
	currentFloor          int
	floorRequestList      []int
	direction             string
	overweight            bool
	completedRequestsList []int
}

func NewElevator(id int, status string, amountOfFloor int, currentFloor int) *Elevator {

	elevator := &Elevator{id: id, status: "idle", amountOfFloor: amountOfFloor, currentFloor: currentFloor}

	return elevator
}

func (e *Elevator) move() {

if (this.floorRequestsList.Count != 0) {
	var destination = this.floorRequestsList[0];
	this.status = "moving";
	if (this.currentFloor < destination) {
		this.direction = "up";
		this.sortFloorList();
		if (this.currentFloor < destination) {
			this.currentFloor++;
			this.screenDisplay = this.currentFloor;
		}
	}
	else if (this.currentFloor > destination) {
		this.direction = "down";
		this.sortFloorList();
		while (this.currentFloor > destination) {
			this.currentFloor--;
			this.screenDisplay = this.currentFloor;
		}
	}
	this.status = "stopped";
	this.operateDoors();
	this.completedRequestsList.Add(this.floorRequestsList[0]);
	this.floorRequestsList.RemoveAt(0);
}
this.status = "idle";

}
