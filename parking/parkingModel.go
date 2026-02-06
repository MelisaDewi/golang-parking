package parking

import (
	"errors"
	"fmt"
	"strconv"
)

var noAvailableSpace = errors.New("no available parking space")
var wrongParkingLot = errors.New("wrong parking lot")
var unrecognizedParkingLot = errors.New("unrecognized parking lot")
var unrecognizedParkingTicket = errors.New("unrecognized parking ticket")
var carAlreadyParked = errors.New("car already parked")
var cannotFindCar = errors.New("car is not in the parking lot")
var noSubscribe = errors.New("does not subscribe")

var ticketNum int = 0

type Car struct {
	Tipe     string
	Colour   string
	PlateNum string
}

type ParkingItf interface {
	AddCar(car Car) (string, error)
	GetCar(ticket string) (string, error)
}

type Parking struct {
	Name       string
	MaxLot     int
	LotCounter int
	Status     bool
	Car        []Car
	Ticket     []Ticket
}

type Attendant struct {
	Name         string
	ParkingLot   []*Parking
	Car          *Car
	Ticket       string
	Subscription Subscription
}

// type Subscription struct {
// 	ParkingLot Parking
// 	Subscribe  bool
// }

type Subscription (map[string]bool)

type Ticket struct {
	Number string
	Car    Car
}

type Tickets (map[string]Car)

func NewCar(tipe, colour, plateNum string) *Car {
	return &Car{
		Tipe:     tipe,
		Colour:   colour,
		PlateNum: plateNum,
	}
}

func NewParking(name string, maxLot int) *Parking {
	return &Parking{
		Name:   name,
		Status: false,
		MaxLot: maxLot,
		// TicketCounter: ticketCounter,
	}
}

func NewTicket(number string, car Car) *Ticket {
	return &Ticket{
		Number: number,
		Car:    car,
	}
}

func NewAttendant(name string, lot *Parking, subscribe Subscription) *Attendant {
	return &Attendant{
		Name:         name,
		ParkingLot:   []*Parking{lot},
		Subscription: subscribe,
	}
}

func (a *Attendant) AddParkingLot(parkir *Parking) {
	a.ParkingLot = append(a.ParkingLot, parkir)
}

func (a *Attendant) ToggleSubscription(parkir Parking) {
	a.Subscription[parkir.Name] = !a.Subscription[parkir.Name]
}

func (a *Attendant) CheckFUll(parkir Parking) (bool, error) {
	_, ok := a.Subscription[parkir.Name]
	if ok {
		if parkir.Status == true {
			return true, noAvailableSpace
		}
	}
	return false, noSubscribe
}

var contoh int

var contoh2 Car

func (p *Parking) AddCar(car *Car) (string, error) {
	// fmt.Println(p.Car)
	for i := range len(p.Car) {
		// fmt.Println(i)
		fmt.Println(len(p.Car))
		fmt.Println("masuk", car.PlateNum, "dicek", p.Car[i].PlateNum)
		if car.PlateNum == p.Car[i].PlateNum {
			return "Car already parked", carAlreadyParked
		}
		// fmt.Println("sampe sini")
	}
	if p.MaxLot > len(p.Car) {
		p.LotCounter++
		ticketNumber := "ticket#" + strconv.Itoa(ticketNum)
		ticketNum++
		fmt.Println("Car yg diappend ", car)
		p.Car = append(p.Car, *car)
		fmt.Println("Car parking lot ", p.Car)
		ticket := NewTicket(ticketNumber, *car)
		p.Ticket = append(p.Ticket, *ticket)
		if p.LotCounter == p.MaxLot {
			p.Status = true
		}
		// p.Ticket[ticket.Number] = p.Ticket[ticket.Car.Tipe]
		return ticket.Number, nil
	} else {
		return "Cannot park", noAvailableSpace
	}
}

func (p *Parking) GetCar(ticket string) (string, error) {
	if len(ticket) < 7 {
		return "Invalid ticket", unrecognizedParkingTicket
	}
	if ticket[:7] != "ticket#" {
		return "Invalid ticket", unrecognizedParkingTicket
	}
	for i := range p.Ticket {
		// fmt.Println("Avail parking tickets ", p.Ticket)
		// fmt.Println("Avail cars ", p.Car)
		// fmt.Println("Parking ticket ", p.Ticket[i].Number)
		// fmt.Println("Yg dipake ", ticket)
		if p.Ticket[i].Number == ticket {
			for j := range len(p.Car) {
				fmt.Println(p.Car[j].PlateNum)
				fmt.Println(p.Ticket[j].Car.PlateNum)
				if p.Car[j].PlateNum == p.Ticket[i].Car.PlateNum {
					p.Car = append(p.Car[:j], p.Car[j+1:]...)
					p.Ticket = append(p.Ticket[:i], p.Ticket[i+1:]...)
					// fmt.Println("Pas udh diambil tiketnya ", p.Ticket)
					// fmt.Println("Pas udh diambil mobilnya ", p.Car)

					return "Car successfully unparked", nil
				}
			}
		}
	}

	return "Invalid ticket", unrecognizedParkingTicket
}

func (a *Attendant) AttAddCar(car *Car) (string, error) {
	a.Car = car
	for i := range len(a.ParkingLot) {
		p := a.ParkingLot[i]
		val, _ := a.CheckFUll(*p)
		if val == true {
			fmt.Println("Parkiran ", p.Name, " penuh")
			continue
		}
		res, err := p.AddCar(a.Car)
		if err == nil {
			a.Ticket = res
			a.Car = nil
			fmt.Println("Diisi di parkiran ", p.Name, " mobilnya ", p.Car)
			return res, nil
		} else {
			return res, err
		}
	}
	return "Kayaknya parkirannya penuh", noAvailableSpace
}

func (a *Attendant) AttGetCar(ticket string) (string, error) {
	// if len(ticket) < 7 {
	// 	return "Invalid ticket", unrecognizedParkingTicket
	// }
	// if ticket[:7] != "ticket#" {
	// 	return "Invalid ticket", unrecognizedParkingTicket
	// }

	// for i := range len(a.Ticket) {
	// 	if ticket == a.Ticket[i] {
	for i := range len(a.ParkingLot) {
		p := a.ParkingLot[i]
		res, err := p.GetCar(ticket)
		if err == nil {
			a.Ticket = ""
			return res, nil
		}
	}
	// 	}
	// }

	return "Ga nemu mobilnya", cannotFindCar
}
