package model

import "errors"

var (
	carAlreadyParked          = errors.New("car already parked")
	unrecognizedParkingTicket = errors.New("unrecognized parking ticket")
)

type Car struct {
	Tipe     string
	Colour   string
	PlateNum string
}

// type ParkingItf interface {
// 	AddCar(car Car) (string, error)
// 	GetCar(ticket string) (string, error)
// }

type ParkingSystem struct {
	CarNum map[string]struct{}
	Ticket map[string]struct{}
}

type ParkingItf interface {
	AddCar(ps *ParkingSystem, car *Car) (string, error)
	GetCar(ps *ParkingSystem, ticket string) (string, error)
	CheckCarExist(car *Car) (string, error)

	GetName() string
	GetMaximum() int
	GetStatus() bool
	GetFreeSpace() int
}

func NewCar(tipe, colour, plateNum string) *Car {
	return &Car{
		Tipe:     tipe,
		Colour:   colour,
		PlateNum: plateNum,
	}
}

func NewParkingSystem() ParkingSystem {
	return ParkingSystem{
		CarNum: make(map[string]struct{}),
		Ticket: make(map[string]struct{}),
		// TicketCounter: ticketCounter,
	}
}

func (ps *ParkingSystem) CheckCarExist(car *Car) (string, error) {
	_, ok := ps.CarNum[car.PlateNum]
	if ok {
		return "Car already parked", carAlreadyParked
	}
	return "Car is not recognized", nil
}

func (ps *ParkingSystem) CheckTicketExist(ticket string) (string, error) {
	_, ok := ps.Ticket[ticket]
	if ok {
		return "Ticket exists", nil
	}
	return "Invalid ticket", unrecognizedParkingTicket
}
