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

	parking.CliParking()

	// var nyobaSub map[string]bool{"test":false}
	// var mapBool map[string]bool = make(map[string]bool)
	// var subscription parking.Subscription
	// parkingSystem := model.NewParkingSystem()
	// // parkir, _ := parking.NewParking("parkir 1", 1)
	// // parkir2, _ := parking.NewParking("parkir 2", 2)
	// // parkir3, _ := parking.NewParking("parkir 3", 3)
	// // parkirs := []parking.Parking{*parkir, *parkir2, *parkir3}
	// // for i := range len(parkirs) {
	// // 	mapBool[parkirs[i].Name] = false
	// // }

	// fs := parking.SortFreeSpace{}
	// ml := parking.SortMaxLot{}
	// sq := parking.SortSequential{}

	// // Sorting choices : default, sortMaxLot, sortFreeSpace
	// attendant := parking.NewAttendantNoLot("nama si attendant", sq)
	// attendant.AddParkingLot(parkir2, parkir3)

	// attendant2 := parking.NewAttendant("nama si attendant kedua", parkir, fs)
	// attendant2.AddParkingLot(parkir2, parkir3)

	// parkir.Register(attendant)
	// parkir2.Register(attendant)
	// parkir3.Register(attendant)

	// parkir.Register(attendant2)

	// fmt.Println(attendant.ParkingLot)
	// fmt.Println(attendant.ParkingLot[0].GetName())
	// fmt.Println(attendant.ParkingLot[1].GetName())
	// fmt.Println(attendant2.ParkingLot)

	// attendant.ArrangeParkingLot()
	// attendant2.ArrangeParkingLot()

	// fmt.Println(attendant.ParkingLot[0].GetName())
	// fmt.Println(attendant.ParkingLot[1].GetName())
	// fmt.Println(attendant2.ParkingLot)
	// attendants := []parking.Attendant{*attendant, *attendant2}

	/*
		mauApa := ""
		parkingLotCounter := 1

		for {
			mauApa = io("\n1 -> Setup\n2 -> Park\n3 -> Un-park\n4 -> Status\n5 -> Change Strategy\n6 -> Exit\n\nInput menu > ")
			if mauApa == "1" {
				parkingname := "Parking " + strconv.Itoa(parkingLotCounter)
				parkingcapacity := io("Capacity > ")
				parkingcapacityint, _ := strconv.Atoi(parkingcapacity)

				parkir, _ := parking.NewParking(parkingname, parkingcapacityint)
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

	*/

	// mobilt := io("Isi tipe mobil: ")
	// mobilc := io("Isi warna mobil: ")
	// mobilp := io("Isi plat nomor mobil: ")

	// mobil1 := parking.NewCar(mobilt, mobilc, mobilp)

	// mobil1 := model.NewCar("tipe1", "koneng", "0000")
	// mobil2 := model.NewCar("tipe2", "putih", "1244")
	// mobil3 := model.NewCar("tipe3", "hitam", "4444")
	// mobil4 := model.NewCar("tipe4", "biru", "7890")
	// mobil5 := model.NewCar("tipe5", "merah", "87654")
	// mobil6 := model.NewCar("tipe6", "poleng", "234566")
	// mobil7 := model.NewCar("tipe7", "ga tau lagi", "7777777")

	// fmt.Println("Add car ke 1")
	// ticket, err := attendant.AddCar(&parkingSystem, mobil1)

	// if err == nil {
	// 	fmt.Println("tiket mobil 1: ", ticket)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("--------------------------------------")
	// fmt.Println("Add car ke 2")
	// ticket2, err := attendant.AddCar(&parkingSystem, mobil2)

	// if err == nil {
	// 	fmt.Println("tiket mobil 2: ", ticket2)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("--------------------------------------")
	// attendant.ChangeParkingType(fs)
	// fmt.Println("Add car ke 3")
	// ticket3, err := attendant.AddCar(&parkingSystem, mobil3)

	// if err == nil {
	// 	fmt.Println("tiket mobil 3: ", ticket3)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("--------------------------------------")
	// fmt.Println("Nyoba ngambil mobil tiket valid")
	// // fmt.Println(parkir.Car)
	// // fmt.Println(ticket)
	// ngambil, err := attendant.GetCar(&parkingSystem, ticket)
	// if err == nil {
	// 	fmt.Println(ngambil)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// // fmt.Println(parkir.Car)
	// fmt.Println("--------------------------------------")
	// fmt.Println("Add car ke 4")
	// ticket4, err := attendant2.AddCar(&parkingSystem, mobil4)

	// if err == nil {
	// 	fmt.Println("tiket mobil 4: ", ticket4)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("--------------------------------------")
	// fmt.Println("Add car ke 5")
	// ticket5, err := attendant2.AddCar(&parkingSystem, mobil5)

	// if err == nil {
	// 	fmt.Println("tiket mobil 5: ", ticket5)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("--------------------------------------")
	// fmt.Println("Add car ke 6")
	// ticket6, err := attendant2.AddCar(&parkingSystem, mobil6)

	// if err == nil {
	// 	fmt.Println("tiket mobil 6: ", ticket6)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("--------------------------------------")
	// fmt.Println("Add car ke 7")
	// ticket7, err := attendant2.AddCar(&parkingSystem, mobil7)

	// if err == nil {
	// 	fmt.Println("tiket mobil 7: ", ticket7)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("--------------------------------------")

	// /*fmt.Println("Nyoba isi di parking 3")

	// ticket, err := attendant.AddCar(&parkingSystem, mobil1)

	// if err == nil {
	// 	fmt.Println("tiket mobil 1: ", ticket)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// // fmt.Println("Parking lot yg di main ", parkir.Car)
	// fmt.Println("----------------------------")
	// fmt.Println("Nyoba isi yang sama")
	// // fmt.Println(ticket1)
	// // fmt.Println(parkir.Ticket)
	// // fmt.Println(parkir.TicketCounter)
	// ticket2, err2 := attendant.AddCar(&parkingSystem, mobil1)
	// if err2 == nil {
	// 	fmt.Println(ticket2)
	// } else {
	// 	fmt.Println(err2.Error())
	// }
	// fmt.Println("----------------------------")
	// fmt.Println("Nyoba isi di parking 3")

	// ticket1, err1 := attendant.AddCar(&parkingSystem, mobil2)

	// if err1 == nil {
	// 	fmt.Println("tiket mobil 2: ", ticket1)
	// } else {
	// 	fmt.Println(err1.Error())
	// }
	// fmt.Println(parkir3.Car)
	// // fmt.Println(parkir.Car)
	// fmt.Println("----------------------------")
	// fmt.Println("Nyoba isi parkir 1")
	// // res, err := attendant.CheckFUll(*parkir)
	// // if res == true {
	// // 	fmt.Println("Parkiran 1 penuh")
	// // }

	// ticket3, err3 := attendant.AddCar(&parkingSystem, mobil3)
	// if err3 == nil {
	// 	fmt.Println(ticket3)
	// } else {
	// 	fmt.Println(err3.Error())
	// }
	// fmt.Println(parkir.Car)

	// fmt.Println("----------------------------")
	// fmt.Println("Nyoba di attendant 2")

	// ticket4, err4 := attendant2.AddCar(&parkingSystem, mobil3)
	// if err4 == nil {
	// 	fmt.Println(ticket4)
	// } else {
	// 	fmt.Println(err4.Error())
	// }
	// // fmt.Println(parkir.Car)

	// fmt.Println("----------------------------")
	// fmt.Println("Nyoba ngambil mobil tiket valid")
	// // fmt.Println(parkir.Car)
	// // fmt.Println(ticket)
	// ngambil, err := attendant.GetCar(&parkingSystem, ticket1)
	// if err == nil {
	// 	fmt.Println(ngambil)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// // fmt.Println(parkir.Car)
	// fmt.Println("----------------------------")
	// fmt.Println("Nyoba ngambil mobil tiket ga valid punya sebelah")
	// ngambil2, err := attendant2.GetCar(&parkingSystem, ticket3)
	// if err == nil {
	// 	fmt.Println(ngambil2)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println("----------------------------")
	// fmt.Println("Nyoba ngambil mobil tiket valid")
	// ngambil3, err := attendant.GetCar(&parkingSystem, ticket)
	// if err == nil {
	// 	fmt.Println(ngambil3)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println("----------------------------")
	// fmt.Println("Nyoba ngambil mobil tiket valid")

	// ngambil4, err := attendant2.GetCar(&parkingSystem, ticket4)
	// if err == nil {
	// 	fmt.Println(ngambil4)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// attendant.ArrangeParkingLotMaxLot()
	// */

}
