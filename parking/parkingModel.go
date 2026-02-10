package parking

import (
	"errors"
	"fmt"
	"parkingLot/model"
	"parkingLot/observer"
	"sort"
	"strconv"
)

var (
	maxLotLessThanZero        = errors.New("max lot should be at least 0")
	noAvailableSpace          = errors.New("no available parking space")
	wrongParkingLot           = errors.New("wrong parking lot")
	unrecognizedParkingLot    = errors.New("unrecognized parking lot")
	unrecognizedParkingTicket = errors.New("unrecognized parking ticket")
	carAlreadyParked          = errors.New("car already parked")
	cannotFindCar             = errors.New("car is not in the parking lot")
	noSubscribe               = errors.New("does not subscribe")
)

var (
	ticketNum   int = 0
	attendantId int = 0
)

type ParkingStyle int

const (
	Sequential ParkingStyle = iota
	HighestMaxLot
	HighestFreeSpace
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

type Parking struct {
	Name             string
	MaxLot           int
	LotCounter       int
	FreeSpaceCounter int
	Status           bool
	Car              []model.Car
	Ticket           []Ticket
	observerList     []observer.Observer
}

type Attendant struct {
	id                      string
	Name                    string
	ParkingLot              []model.ParkingItf
	ParkingLotSort          []model.ParkingItf
	ParkingLotSortFreeSpace []model.ParkingItf
	Car                     *model.Car
	Ticket                  string
	ParkirFull              ParkirFull
	styleSort               ParkingStyle
	// Subscription Subscription
}

// type Subscription struct {
// 	ParkingLot Parking
// 	Subscribe  bool
// }

type ParkirFull (map[string]bool)

type Subscription (map[string]bool)

type Ticket struct {
	Number string
	Car    model.Car
}

type Tickets (map[string]Car)

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
func NewParking(name string, maxLot int) (*Parking, error) {
	if maxLot < 0 {
		return &Parking{
			Name:             name,
			Status:           false,
			MaxLot:           0,
			FreeSpaceCounter: 0,
			LotCounter:       0,
			// TicketCounter: ticketCounter,
		}, maxLotLessThanZero
	}
	return &Parking{
		Name:             name,
		Status:           false,
		MaxLot:           maxLot,
		FreeSpaceCounter: maxLot,
		LotCounter:       0,
		// TicketCounter: ticketCounter,
	}, nil
}

func NewTicket(number string, car model.Car) *Ticket {
	return &Ticket{
		Number: number,
		Car:    car,
	}
}

func NewAttendant(name string, lot model.ParkingItf, style ParkingStyle) *Attendant {
	attendantId++
	return &Attendant{
		id:                      string(rune(attendantId)),
		Name:                    name,
		ParkingLot:              []model.ParkingItf{lot},
		ParkingLotSort:          []model.ParkingItf{lot},
		ParkingLotSortFreeSpace: []model.ParkingItf{lot},
		ParkirFull:              ParkirFull{},
		styleSort:               style,
		// Subscription: subscribe,
	}
}

func (p *Parking) GetName() string {
	return p.Name
}

func (p *Parking) GetMaximum() int {
	return p.MaxLot
}

func (p *Parking) GetStatus() bool {
	return p.Status
}

func (p *Parking) GetFreeSpace() int {
	return p.FreeSpaceCounter
}

func (a *Attendant) AddParkingLot(parkir ...model.ParkingItf) {
	a.ParkingLot = append(a.ParkingLot, parkir...)
	a.ParkingLotSort = append(a.ParkingLotSort, parkir...)
	a.ParkingLotSortFreeSpace = append(a.ParkingLotSort, parkir...)
	a.ArrangeParkingLotMaxLot()
	a.ArrangeParkingLotFreeSpace()
	// fmt.Println("Parking lot free space", a.ParkingLotSortFreeSpace)

}

func (a *Attendant) Update(name string, status bool) bool {
	if status {
		fmt.Println("Dikirim ke attendant", a.Name, "Parking ", name, " full")
		a.ParkirFull[name] = status
		return status
	} else {
		fmt.Println("Dikirim ke attendant", a.Name, "Parking ", name, " not full")
		a.ParkirFull[name] = status
		return status
	}
}

func (a *Attendant) GetID() string {
	return a.id
}

func (a *Attendant) ArrangeParkingLotMaxLot() {
	sort.Slice(a.ParkingLotSort, func(i, j int) bool {
		return a.ParkingLotSort[i].GetMaximum() > a.ParkingLotSort[j].GetMaximum()
	})
	// for _, v := range a.ParkingLotSort {
	// 	fmt.Println(v)
	// }
}

func (a *Attendant) ArrangeParkingLotFreeSpace() {
	sort.Slice(a.ParkingLotSortFreeSpace, func(i, j int) bool {
		return a.ParkingLotSortFreeSpace[i].GetFreeSpace() > a.ParkingLotSortFreeSpace[j].GetFreeSpace()
	})
	// for _, v := range a.ParkingLotSortFreeSpace {
	// 	fmt.Println("Function sorting free space")
	// 	fmt.Println(v.GetName())
	// }
}

func (a *Attendant) ChangeStyle(style ParkingStyle) {
	a.styleSort = style
}

// func (a *Attendant) ToggleSubscription(parkir Parking) {
// 	a.Subscription[parkir.Name] = !a.Subscription[parkir.Name]
// }

// func (a *Attendant) CheckFUll(parkir Parking) (bool, error) {
// 	_, ok := a.Subscription[parkir.Name]
// 	if ok {
// 		if parkir.Status == true {
// 			return true, noAvailableSpace
// 		}
// 	}
// 	return false, noSubscribe
// }

func (p *Parking) Register(o observer.Observer) {
	p.observerList = append(p.observerList, o)
}

func (p *Parking) Deregister(o observer.Observer) {
	p.observerList = removeFromslice(p.observerList, o)
}

func (p *Parking) NotifyAll() {
	for _, observer := range p.observerList {
		observer.Update(p.Name, p.Status)
	}
}

func removeFromslice(observerList []observer.Observer, observerToRemove observer.Observer) []observer.Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

func (p *Parking) CheckFull() bool {
	fmt.Println(p.MaxLot-1, p.LotCounter)
	if p.MaxLot == p.LotCounter {
		p.Status = true
		p.NotifyAll()
		return true
	} else if p.MaxLot-1 == p.LotCounter {
		p.Status = false
		p.NotifyAll()
		return false
	}
	return false
}

func (p *Parking) CheckCarExist(car *model.Car) (string, error) {
	for i := range len(p.Car) {
		// fmt.Println(i)
		// fmt.Println(len(p.Car))
		// fmt.Println("masuk", car.PlateNum, "dicek", p.Car[i].PlateNum)
		if car.PlateNum == p.Car[i].PlateNum {
			return "Car already parked", carAlreadyParked
		}
		// fmt.Println("sampe sini")
	}
	return "Car is not recognized", nil
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

func (p *Parking) AddCar(ps *model.ParkingSystem, car *model.Car) (string, error) {
	// fmt.Println(p.Car)
	// for i := range len(p.Car) {
	// 	// fmt.Println(i)
	// 	fmt.Println(len(p.Car))
	// 	fmt.Println("masuk", car.PlateNum, "dicek", p.Car[i].PlateNum)
	// 	if car.PlateNum == p.Car[i].PlateNum {
	// 		return "Car already parked", carAlreadyParked
	// 	}
	// 	// fmt.Println("sampe sini")
	// }
	if p.MaxLot > len(p.Car) {
		p.LotCounter++
		p.FreeSpaceCounter--
		ticketNumber := "ticket#" + strconv.Itoa(ticketNum)
		ticketNum++
		// fmt.Println("Car yg diappend ", car)
		p.Car = append(p.Car, *car)
		// fmt.Println("Car parking lot ", p.Car)
		ticket := NewTicket(ticketNumber, *car)
		p.Ticket = append(p.Ticket, *ticket)
		ps.Ticket[ticketNumber] = struct{}{}
		ps.CarNum[car.PlateNum] = struct{}{}
		if p.MaxLot == p.LotCounter {

			p.CheckFull()
		}
		// p.Ticket[ticket.Number] = p.Ticket[ticket.Car.Tipe]
		return ticket.Number, nil
	} else {
		// p.Status = true
		// p.NotifyAll()
		return "Cannot park", noAvailableSpace
	}
}

func (p *Parking) GetCar(ps *model.ParkingSystem, ticket string) (string, error) {
	// fmt.Println(ticket)
	_, err := ps.CheckTicketExist(ticket)
	if err == nil {
		for i := range p.Ticket {
			// fmt.Println("Avail parking tickets ", p.Ticket)
			// fmt.Println("Avail cars ", p.Car)
			// fmt.Println("Parking ticket ", p.Ticket[i].Number)
			// fmt.Println("Yg dipake ", ticket)
			if p.Ticket[i].Number == ticket {
				for j := range len(p.Car) {
					// fmt.Println(p.Car[j].PlateNum)
					// fmt.Println(p.Ticket[j].Car.PlateNum)
					if p.Car[j].PlateNum == p.Ticket[i].Car.PlateNum {
						p.Car = append(p.Car[:j], p.Car[j+1:]...)
						p.Ticket = append(p.Ticket[:i], p.Ticket[i+1:]...)
						p.LotCounter--
						p.FreeSpaceCounter++
						// fmt.Println(p.LotCounter, p.MaxLot-1)
						if p.LotCounter == p.MaxLot-1 {
							p.CheckFull()
						}
						// fmt.Println("Pas udh diambil tiketnya ", p.Ticket)
						// fmt.Println("Pas udh diambil mobilnya ", p.Car)

						return "Car successfully unparked", nil
					}
				}
			}
		}
	}
	// if len(ticket) < 7 {
	// 	return "Invalid ticket", unrecognizedParkingTicket
	// }
	// if ticket[:7] != "ticket#" {
	// 	return "Invalid ticket", unrecognizedParkingTicket
	// }

	return "Invalid ticket", unrecognizedParkingTicket
}

func (a *Attendant) AddCar(ps *model.ParkingSystem, car *model.Car) (string, error) {
	a.Car = car
	parkingUsed := a.ParkingLot
	if a.styleSort == HighestMaxLot {
		parkingUsed = a.ParkingLotSort
	} else if a.styleSort == HighestFreeSpace {
		a.ArrangeParkingLotFreeSpace()
		parkingUsed = a.ParkingLotSortFreeSpace
	}
	res, err := ps.CheckCarExist(car)
	if err != nil {
		return res, err
	}

	// fmt.Println("Parking used ", parkingUsed[len(parkingUsed)-1].GetName())
	// fmt.Println("Len i", len(parkingUsed))
	for i := range len(parkingUsed) {
		p := parkingUsed[i]
		// if p.GetStatus() {
		// 	continue
		// }
		// res, err := p.CheckCarExist(car)
		// if err != nil {
		// 	return res, err
		// }
		// fmt.Println("nama parkir yg dimasukin ", p.GetName())
		// fmt.Println("Parkir yg dicek sebelum notif", p.GetName())
		if a.ParkirFull[p.GetName()] {
			continue
		}
		// fmt.Println("Parkir yg dicek", p.GetName())
		// val, _ := a.CheckFUll(*p)
		// if val == true {
		// 	fmt.Println("Parkiran ", p.Name, " penuh")
		// 	continue
		// }
		res, err = p.AddCar(ps, a.Car)
		if err == nil {
			a.Ticket = res
			a.Car = nil
			fmt.Println("Diisi di parkiran ", p.GetName(), " mobilnya ", car)
			// p.CheckFull()
			return res, nil
		}
	}
	return "Kayaknya parkirannya penuh", noAvailableSpace
}

func (a *Attendant) GetCar(ps *model.ParkingSystem, ticket string) (string, error) {
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
		res, err := p.GetCar(ps, ticket)
		if err == nil {
			a.Ticket = ""
			return res, nil
		}
	}
	// 	}
	// }

	return "Ga nemu mobilnya", cannotFindCar
}
