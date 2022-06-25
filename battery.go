package main

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
	servedFloor               Column
	elevator                  Elevator
	amountOfElevators         int
	servedFloorList           []int
}

func NewBattery(id, amountOfColumns, amountOfFloors, amountOfBasements, amountOfElevatorPerColumn int) *Battery {

	battery := &Battery{id: id, amountOfColumns: amountOfColumns, amountOfFloors: amountOfFloors, amountOfElevatorPerColumn: amountOfElevatorPerColumn}

	return battery

	// 	this.ID = _ID;
	// 	this.status = "online";
	// 	this.columnsList = new List<Column>();
	// 	this.floorRequestButtonList = new List<FloorRequestButton>();

	// 	if (_amountOfBasements > 0) {
	// 		this.createBasementFloorRequestButtons(_amountOfBasements);
	// 		this.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn);
	// 		_amountOfColumns--;
	// 	}

	// 	this.createFloorRequestButtons(_amountOfFloors);
	// 	this.createColumns(_amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn);

}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {

	bestColumn := b.columnsList[0]

	for _, column := range b.columnsList {
		for _, x := range column.servedFloor {
			if x == _requestedFloor {
				return &column
			}
		}
	}
	return &bestColumn
}

// //Simulate when a user press a button at the lobby
// func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {

// }
