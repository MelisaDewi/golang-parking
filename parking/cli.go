package parking

import (
	"bufio"
	"fmt"
	"os"
	"parkingLot/model"
	"strconv"
)

func io(textInput string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(textInput)
	text, _ := reader.ReadString('\n')
	// fmt.Printf("%t, %v, %v", text, text, len(text))
	return text[:len(text)-1]
}

func ioInt(textInput string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(textInput)
	text, _ := reader.ReadString('\n')
	textInt, _ := strconv.Atoi(text)
	// fmt.Printf("%t, %v, %v", text, text, len(text))
	return textInt
}

func CliParking() {
	mauApa := ""
	parkingLotCounter := 1
	parkingSystem := model.NewParkingSystem()

	fs := SortFreeSpace{}
	ml := SortMaxLot{}
	sq := SortSequential{}

	// Sorting choices : default, sortMaxLot, sortFreeSpace
	attendant := NewAttendantNoLot("nama si attendant", sq)

	for {
		mauApa = io("\n1 -> Setup\n2 -> Park\n3 -> Un-park\n4 -> Status\n5 -> Change Strategy\n6 -> Exit\n\nInput menu > ")
		if mauApa == "1" {
			parkingname := "Parking " + strconv.Itoa(parkingLotCounter)
			parkingcapacity := io("Capacity > ")
			parkingcapacityint, _ := strconv.Atoi(parkingcapacity)

			parkir, _ := NewParking(parkingname, parkingcapacityint)
			// parkirs = append(parkirs, *parkir)
			parkingLotCounter++
			attendant.AddParkingLot(parkir)

		} else if mauApa == "2" {
			if len(attendant.ParkingLot) < 1 {
				fmt.Println("No parking to park into")
				continue
			}
			attendant.CheckParkingExist()
			mobilp := io("Plate number > ")
			mobil := model.NewCar("tipe mobil", "warna mobil", mobilp)
			ticket, err := attendant.AddCar(&parkingSystem, mobil)
			if err != nil {
				fmt.Println("Error:", err.Error())
				continue
			}
			// tickets = append(tickets, ticket)
			fmt.Println("Your ticket ID : ", ticket)

		} else if mauApa == "3" {
			// fmt.Println("tiket di parkir ", parkir.Ticket)
			tickets, err := parkingSystem.AvailableTickets()
			if err != nil {
				fmt.Println("Ga ada tiket bro")
			} else {
				fmt.Println("tiket di tangan", tickets)
				dikasi := io("Tiket ID > ")
				ngambil, err := attendant.GetCar(&parkingSystem, dikasi)
				if err == nil {
					fmt.Println(ngambil)
					// for j := range len(tickets) {
					// 	if tickets[j] == dikasi {
					// 		tickets = append(tickets[:j], tickets[j+1:]...)
					// 	}
					// }
				} else {
					fmt.Println(err.Error())
				}
			}
		} else if mauApa == "4" {
			if len(attendant.ParkingLot) < 1 {
				fmt.Println("No parking to get info")
				continue
			}
			for i := range len(attendant.ParkingLot) {
				fmt.Println(attendant.ParkingLot[i].GetName(), ":", "\n  - Capacity:", attendant.ParkingLot[i].GetMaximum(), "\n  - Occupied:", attendant.ParkingLot[i].GetOccupiedSpace())
			}
		} else if mauApa == "5" {
			style := ""
			switch attendant.ParkingStyle {
			case sq:
				style = "1 (Sequential)"
			case ml:
				style = "2 (Most Capacity)"
			case fs:
				style = "3 (Most Available)"
			}
			parkingT := io(fmt.Sprint("1 -> Sequential\n2 -> Most Capacity\n3 -> Most Available\nCurrent: ", style, ", Change to > "))
			parkingTint, _ := strconv.Atoi(parkingT)
			switch parkingTint {
			case 1:
				attendant.ChangeParkingType(sq)
				fmt.Println("Changed attendant strategy to sequential")
			case 2:
				attendant.ChangeParkingType(ml)
				fmt.Println("Changed attendant strategy to most capacity")
			case 3:
				attendant.ChangeParkingType(fs)
				fmt.Println("Changed attendant strategy to most available")
			default:
				fmt.Println("Isi 1-3 cok")
			}
		} else if mauApa == "6" {
			fmt.Println("Good bye!")
			break
		} else {
			fmt.Println("Isi 1-6 cuy")
		}
	}
}
