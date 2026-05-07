package main

import (
	"bufio"
	"fmt"
	"os"
	"parkingLot/parking"
)

func io(textInput string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(textInput)
	text, _ := reader.ReadString('\n')
	// fmt.Printf("%t, %v, %v", text, text, len(text))
	return text[:len(text)-1]
}

func main() {
	// mauApa := ""
	// parkingLotCounter := 1
	// parkingSystem := model.NewParkingSystem()

	// fs := SortFreeSpace{}
	// ml := SortMaxLot{}
	// sq := SortSequential{}

	// // Sorting choices : default, sortMaxLot, sortFreeSpace
	// attendant := NewAttendantNoLot("nama si attendant", sq)

	// parkingSystem.Register(attendant)

	parking.CliParking()

}
