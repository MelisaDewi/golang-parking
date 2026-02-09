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

	// var nyobaSub map[string]bool{"test":false}
	// var mapBool map[string]bool = make(map[string]bool)
	// var subscription parking.Subscription
	parkingSystem := parking.NewParkingSystem()
	parkir := parking.NewParking("parkir 1", 1)
	parkir2 := parking.NewParking("parkir 2", 1)
	parkir3 := parking.NewParking("parkir 3", 1)
	// parkirs := []parking.Parking{*parkir, *parkir2, *parkir3}
	// for i := range len(parkirs) {
	// 	mapBool[parkirs[i].Name] = false
	// }
	attendant := parking.NewAttendant("nama si attendant", parkir)
	attendant.AddParkingLot(parkir3)

	attendant2 := parking.NewAttendant("nama si attendant kedua", parkir2)
	attendant2.AddParkingLot(parkir3)

	parkir.Register(attendant)
	parkir3.Register(attendant)

	parkir2.Register(attendant2)
	// attendants := []parking.Attendant{*attendant, *attendant2}

	// mauApa := ""

	// lanjut := true

	// mobils := []parking.Car{}
	// tickets := []string{}
	// // var err error7

	// for lanjut {

	// 	mauApa = io("1 -> Bikin mobil baru\n2 -> Masukin mobil\n3 -> Ngambil mobil\n4 -> Exit\n")

	// 	if mauApa == "1" {
	// 		mobilt := io("Isi tipe mobil: ")
	// 		mobilc := io("Isi warna mobil: ")
	// 		mobilp := io("Isi plat nomor mobil: ")

	// 		mobil := parking.NewCar(mobilt, mobilc, mobilp)
	// 		mobils = append(mobils, *mobil)
	// 		fmt.Println(mobils)
	// 	} else if mauApa == "2" {
	// 		if len(mobils) < 1 {
	// 			fmt.Println("Ga ada mobil")
	// 		} else {
	// 			fmt.Println(mobils)
	// 			inputMobil := io("Pilih mobil no brp?")
	// 			mobilPilihanInt, _ := strconv.Atoi(inputMobil)
	// 			ticket, err := attendant.AttAddCar(parkir, &mobils[mobilPilihanInt])
	// 			if err == nil {
	// 				tickets = append(tickets, ticket)
	// 				mobils = append(mobils[:mobilPilihanInt], mobils[mobilPilihanInt+1:]...)
	// 			} else {
	// 				fmt.Println(err.Error())
	// 			}
	// 			fmt.Println(tickets)
	// 		}
	// 	} else if mauApa == "3" {
	// 		// fmt.Println("tiket di parkir ", parkir.Ticket)
	// 		fmt.Println("tiket di tangan", tickets)
	// 		if len(tickets) < 1 {
	// 			fmt.Println("Ga ada tiket bro")
	// 		} else {
	// 			dikasi := io("Tiket mana: ")
	// 			ngambil, err := attendant.AttGetCar(parkir, dikasi)
	// 			if err == nil {
	// 				fmt.Println(ngambil)
	// 				tickets = tickets[1:]
	// 			} else {
	// 				fmt.Println(err.Error())
	// 			}
	// 		}
	// 	} else if mauApa == "4" {
	// 		lanjut = false
	// 	} else {
	// 		fmt.Println("Isi 1-4 cuy")
	// 	}
	// }

	// mobilt := io("Isi tipe mobil: ")
	// mobilc := io("Isi warna mobil: ")
	// mobilp := io("Isi plat nomor mobil: ")

	// mobil1 := parking.NewCar(mobilt, mobilc, mobilp)

	mobil1 := parking.NewCar("tipe1", "koneng", "0000")
	mobil2 := parking.NewCar("tipe2", "putih", "1244")
	mobil3 := parking.NewCar("tipe3", "hitam", "4444")

	fmt.Println("Nyoba isi di parking 1")

	ticket, err := attendant.AttAddCar(&parkingSystem, mobil1)

	if err == nil {
		fmt.Println("tiket mobil 1: ", ticket)
	} else {
		fmt.Println(err.Error())
	}
	// fmt.Println("Parking lot yg di main ", parkir.Car)
	fmt.Println("----------------------------")
	fmt.Println("Nyoba isi yang sama")
	// fmt.Println(ticket1)
	// fmt.Println(parkir.Ticket)
	// fmt.Println(parkir.TicketCounter)
	ticket2, err2 := attendant.AttAddCar(&parkingSystem, mobil1)
	if err2 == nil {
		fmt.Println(ticket2)
	} else {
		fmt.Println(err2.Error())
	}
	fmt.Println("----------------------------")
	fmt.Println("Nyoba isi di parking 3")

	ticket1, err1 := attendant.AttAddCar(&parkingSystem, mobil2)

	if err1 == nil {
		fmt.Println("tiket mobil 2: ", ticket1)
	} else {
		fmt.Println(err1.Error())
	}
	fmt.Println(parkir3.Car)
	// fmt.Println(parkir.Car)
	fmt.Println("----------------------------")
	fmt.Println("Nyoba full")
	// res, err := attendant.CheckFUll(*parkir)
	// if res == true {
	// 	fmt.Println("Parkiran 1 penuh")
	// }

	ticket3, err3 := attendant.AttAddCar(&parkingSystem, mobil3)
	if err3 == nil {
		fmt.Println(ticket3)
	} else {
		fmt.Println(err3.Error())
	}
	fmt.Println(parkir.Car)

	fmt.Println("----------------------------")
	fmt.Println("Nyoba di attendant 2")

	ticket4, err4 := attendant2.AttAddCar(&parkingSystem, mobil3)
	if err4 == nil {
		fmt.Println(ticket4)
	} else {
		fmt.Println(err4.Error())
	}
	// fmt.Println(parkir.Car)

	fmt.Println("----------------------------")
	fmt.Println("Nyoba ngambil mobil tiket valid")
	// fmt.Println(parkir.Car)
	// fmt.Println(ticket)
	ngambil, err := attendant.AttGetCar(&parkingSystem, ticket1)
	if err == nil {
		fmt.Println(ngambil)
	} else {
		fmt.Println(err.Error())
	}
	// fmt.Println(parkir.Car)
	fmt.Println("----------------------------")
	fmt.Println("Nyoba ngambil mobil tiket ga valid punya sebelah")
	ngambil2, err := attendant2.AttGetCar(&parkingSystem, ticket)
	if err == nil {
		fmt.Println(ngambil2)
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println("----------------------------")
	fmt.Println("Nyoba ngambil mobil tiket valid")
	ngambil3, err := attendant.AttGetCar(&parkingSystem, ticket)
	if err == nil {
		fmt.Println(ngambil3)
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println("----------------------------")
	fmt.Println("Nyoba ngambil mobil tiket valid")

	ngambil4, err := attendant2.AttGetCar(&parkingSystem, ticket4)
	if err == nil {
		fmt.Println(ngambil4)
	} else {
		fmt.Println(err.Error())
	}

}
