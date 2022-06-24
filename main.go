package main

import (
	"fmt"
	//"os"
	//"strconv"
)

func main() {

	door := NewDoor(1, "closed")
	fmt.Println(door)
	FloorRequestButton := NewFloorRequestButton(1, "closed", 2, "up")
	fmt.Println(FloorRequestButton)
	CallButton := NewCallButton(1, "moving", 3, "up")
	fmt.Println(CallButton)
	Elevator := NewElevator(1, "idle", 5, 3)
	fmt.Println(Elevator)
	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err == nil {
	// 	runScenario(scenarioNumber)
	// } else {
	// 	fmt.Println(err)
	// }
}
