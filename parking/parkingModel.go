package parking

import (
	"errors"
	"fmt"
	"strconv"
)

var noAvailableSpace = errors.New("no available parking space")
var unrecognizedParkingTicket = errors.New("unrecognized parking ticket")
var carAlreadyParked = errors.New("car already parked")

type Car struct {
	Tipe     string
	Colour   string
	PlateNum string
}

type Parking struct {
	Name          string
	MaxLot        int
	TicketCounter int
	Car           []Car
	Ticket        []Ticket
}

type Ticket struct {
	Number string
	Car    Car
}

func NewCar(tipe, colour, plateNum string) *Car {
	return &Car{
		Tipe:     tipe,
		Colour:   colour,
		PlateNum: plateNum,
	}
}

func NewParking(name string, maxLot int, ticketCounter int) *Parking {
	return &Parking{
		Name:          name,
		MaxLot:        maxLot,
		TicketCounter: ticketCounter,
	}
}

func NewTicket(number string, car Car) *Ticket {
	return &Ticket{
		Number: number,
		Car:    car,
	}
}

func (p *Parking) AddCar(car Car) (string, error) {
	fmt.Println(p.Car)
	for i := range len(p.Car) {
		// fmt.Println(i)
		// fmt.Println(len(p.Car))
		// fmt.Println("masuk", car.PlateNum, "dicek", p.Car[i].PlateNum)
		if car.PlateNum == p.Car[i].PlateNum {
			return "Car already parked", carAlreadyParked
		}
		// fmt.Println("sampe sini")
	}
	if p.MaxLot > len(p.Car) {
		p.TicketCounter++
		ticketNumber := "ticket#" + strconv.Itoa(p.TicketCounter)
		p.Car = append(p.Car, car)
		ticket := NewTicket(ticketNumber, car)
		p.Ticket = append(p.Ticket, *ticket)
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
	for i := range p.TicketCounter {
		fmt.Println("Parking ticket ", p.Ticket[i].Number)
		fmt.Println("Yg dipake ", ticket)
		if p.Ticket[i].Number == ticket {
			for j := range len(p.Car) {
				fmt.Println(p.Car[j].PlateNum)
				fmt.Println(p.Ticket[j].Car.PlateNum)
				// if p.Car[j].PlateNum == p.Ticket[j].Car.PlateNum {
				p.Car = append(p.Car[:j], p.Car[j+1:]...)
				p.Ticket = append(p.Ticket[:j], p.Ticket[j+1:]...)
				return "Car successfully unparked", nil
				// }
			}
		}
	}

	return "Invalid ticket", unrecognizedParkingTicket
}
