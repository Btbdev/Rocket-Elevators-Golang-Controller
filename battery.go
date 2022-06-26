package main

//Attributes of the battery
type Battery struct {
	id                        int
	status                    string
	amountOfFloors            int
	amountOfColumns           int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	column                    Column
	columnsList               []int
	floorRequestButton        FloorRequestButton
	floorRequestButtonList    []int
	servedFloor               int
	elevator                  Elevator
	amountOfElevators         int
	servedFloorList           []int
}

func NewBattery(id, amountOfColumns, amountOfFloors, amountOfBasements, amountOfElevatorPerColumn int) *Battery {

	battery := &Battery{id: id, amountOfColumns: amountOfColumns, amountOfFloors: amountOfFloors, amountOfElevatorPerColumn: amountOfElevatorPerColumn}
	battery.createBasementFloorRequestButtons(battery.amountOfBasements)
	battery.createBasementColumn(battery.amountOfBasements, battery.amountOfElevatorPerColumn)
	battery.createFloorRequestButtons(battery.amountOfFloors)
	battery.createColumns(battery.amountOfColumns, battery.amountOfFloors, battery.amountOfBasements, battery.amountOfElevatorPerColumn)

	return battery

}

func (b *Battery) createBasementColumn(amountOfBasements int, amountOfElevatorPerColumn int) {

}

func (b *Battery) createColumns(amountOfColumns int, amountOfFloors int, amountOfBasements int, amountOfElevatorPerColumn int) {

}

func (b *Battery) createFloorRequestButtons(amountOfFloors int) {

}

func (b *Battery) createBasementFloorRequestButtons(amountOfFloors int) {

}

// func (b *Battery) findBestColumn(_requestedFloor int) *Column {

// 	bestColumn := b.columnsList[0]

// 	for _, column := range b.columnsList {
// 		for _, x := range column.servedFloor {
// 			if x == _requestedFloor {
// 				return &column
// 			}
// 		}
// 	}
// 	return &bestColumn
// }

// //Simulate when a user press a button at the lobby
// func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {

// }
