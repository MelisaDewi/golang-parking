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
	return text
}

func main() {

	parkir := parking.NewParking("parkir 1", 2, 0)

	mauApa := ""

	lanjut := true

	mobils := []parking.Car{}
	tickets := []string{}
	// var err error

	for lanjut {

		mauApa = io("1 -> Bikin mobil baru\n2 -> Masukin mobil\n3 -> Ngambil mobil\n4 -> Exit\n")

		if mauApa == "1\n" {
			mobilt := io("Isi tipe mobil: ")
			mobilc := io("Isi warna mobil: ")
			mobilp := io("Isi plat nomor mobil: ")

			mobil := parking.NewCar(mobilt, mobilc, mobilp)
			mobils = append(mobils, *mobil)
			fmt.Println(mobils)
		} else if mauApa == "2\n" {
			if len(mobils) < 1 {
				fmt.Println("Ga ada mobil")
			} else {
				ticket, err := parkir.AddCar(mobils[0])
				if err == nil {
					tickets = append(tickets, ticket)
					mobils = mobils[1:]
				} else {
					fmt.Println(err.Error())
				}
				fmt.Println(tickets)
			}
		} else if mauApa == "3\n" {
			fmt.Println("tiket di parkir ", parkir.Ticket)
			fmt.Println("tiket di tangan", tickets)
			if len(tickets) < 1 {
				fmt.Println("Ga ada tiket bro")
			} else {
				dikasi := io("Tiket mana: ")
				ngambil, err := parkir.GetCar(dikasi[:len(dikasi)-1])
				if err == nil {
					fmt.Println(ngambil)
					tickets = tickets[1:]
				} else {
					fmt.Println(err.Error())
				}
			}
		} else if mauApa == "4\n" {
			lanjut = false
		} else {
			fmt.Println("Isi 1-4 cuy")
		}
	}

	// mobilt := io("Isi tipe mobil")
	// mobilc := io("Isi warna mobil")
	// mobilp := io("Isi plat nomor mobil")

	// mobil1 := parking.NewCar(mobilt, mobilc, mobilp)
	// mobil2 := parking.NewCar("tipe2", "putih", "1244")
	// mobil3 := parking.NewCar("tipe1", "hitam", "4444")

	// ticket, err := parkir.AddCar(*mobil1)

	// if err == nil {
	// 	fmt.Println("tiket mobil 1: ", ticket)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(parkir.Car)
	// fmt.Println("----------------------------")

	// ticket1, err1 := parkir.AddCar(*mobil2)

	// if err1 == nil {
	// 	fmt.Println("tiket mobil 2: ", ticket1)
	// } else {
	// 	fmt.Println(err1.Error())
	// }
	// fmt.Println(parkir.Car)
	// fmt.Println("----------------------------")
	// // fmt.Println(ticket1)
	// // fmt.Println(parkir.Ticket)
	// // fmt.Println(parkir.TicketCounter)
	// ticket2, err2 := parkir.AddCar(*mobil1)
	// if err2 == nil {
	// 	fmt.Println(parkir.GetCar(ticket2))
	// } else {
	// 	fmt.Println(err2.Error())
	// }
	// fmt.Println(parkir.Car)
	// fmt.Println("----------------------------")
	// fmt.Println("Nyoba full")

	// ticket3, err3 := parkir.AddCar(*mobil3)
	// if err2 == nil {
	// 	fmt.Println(parkir.GetCar(ticket3))
	// } else {
	// 	fmt.Println(err3.Error())
	// }
	// fmt.Println(parkir.Car)
	// fmt.Println("----------------------------")
	// // fmt.Println(parkir.Car)
	// // fmt.Println(ticket)
	// ngambil, err := parkir.GetCar(ticket)
	// if err == nil {
	// 	fmt.Println(ngambil)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(parkir.Car)
	// fmt.Println("----------------------------")

	// ngambil2, err := parkir.GetCar(ticket1)
	// if err == nil {
	// 	fmt.Println(ngambil2)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// ngambil3, err := parkir.GetCar("bukan")
	// if err == nil {
	// 	fmt.Println(ngambil3)
	// } else {
	// 	fmt.Println(err.Error())
	// }

}
