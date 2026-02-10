package parking

import (
	"fmt"
	mockobserver "parkingLot/mock/mock_observer"
	mockparking "parkingLot/mock/mock_parking"
	"parkingLot/model"
	"parkingLot/observer"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var parkingSystem = model.NewParkingSystem()

// var parking = model.NewParking("parkir 1", 2)
var car = model.NewCar("apa", "putih", "1111")
var car2 = model.NewCar("apa", "putih", "2222")

func TestObserver_WithMock_Update_GivenParkir1FullButNoSub_ShouldReturnTheCorrectResult(t *testing.T) {
	// Given.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	observer := mockobserver.NewMockObserver(ctrl)
	parkir, _ := NewParking("parkir 1", 1)

	observer.EXPECT().Update("parkir 1", true).Return(true).Times(0)

	result, err := parkir.AddCar(&parkingSystem, car)

	// Then.
	assert.Equal(t, "ticket#0", result)
	assert.NoError(t, err)
}
func TestObserver_WithMock_Update_GivenParkir1Full_ShouldReturnTheCorrectResult(t *testing.T) {
	// Given.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	observer := mockobserver.NewMockObserver(ctrl)
	parkir, _ := NewParking("parkir 1", 1)
	parkir.Register(observer)

	observer.EXPECT().Update("parkir 1", true).Return(true).Times(1)

	result, err := parkir.AddCar(&parkingSystem, car)

	// Then.
	assert.Equal(t, "ticket#1", result)
	assert.NoError(t, err)
}

func TestObserver_WithMock_Update_GivenParkir1Avail_ShouldReturnTheCorrectResult(t *testing.T) {
	// Given.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	observer := mockobserver.NewMockObserver(ctrl)
	parkir, _ := NewParking("parkir 1", 1)
	ticket, _ := parkir.AddCar(&parkingSystem, car)
	fmt.Println("ticket no 1 ", ticket)
	parkir.Register(observer)

	observer.EXPECT().Update("parkir 1", false).Return(false).Times(1)

	result, err := parkir.GetCar(&parkingSystem, "ticket#2")

	// Then.
	assert.Equal(t, "Car successfully unparked", result)
	assert.NoError(t, err)
}
func TestObserver_WithMock_Update_GivenParkir1AvailNoSub_ShouldReturnTheCorrectResult(t *testing.T) {
	// Given.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	observer := mockobserver.NewMockObserver(ctrl)
	parkir, _ := NewParking("parkir 1", 1)
	ticket, _ := parkir.AddCar(&parkingSystem, car)
	fmt.Println("ticket ", ticket)
	parkir.AddCar(&parkingSystem, car)

	observer.EXPECT().Update("parkir 1", false).Return(false).Times(0)

	result, err := parkir.GetCar(&parkingSystem, "ticket#3")

	// Then.
	assert.Equal(t, "Car successfully unparked", result)
	assert.NoError(t, err)
}
func TestObserver_WithMock_Update_GivenParkir1Avail_ThenAddCar_ShouldReturnTheCorrectResult(t *testing.T) {
	// Given.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	observer := mockobserver.NewMockObserver(ctrl)
	parkir, _ := NewParking("parkir 1", 1)
	ticket, _ := parkir.AddCar(&parkingSystem, car)
	fmt.Println("ticket ", ticket)
	parkir.AddCar(&parkingSystem, car)
	parkir.GetCar(&parkingSystem, "ticket#4")
	parkir.Register(observer)

	observer.EXPECT().Update("parkir 1", true).Return(true).Times(1)

	result, err := parkir.AddCar(&parkingSystem, car)

	// Then.
	assert.Equal(t, "ticket#5", result)
	assert.NoError(t, err)
}

// func TestObserver_WithMock_Update_GivenParkir1NotFull_ShouldReturnTheCorrectResult(t *testing.T) {
// 	// Given.
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	observer := mockobserver.NewMockObserver(ctrl)
// 	// observer := &Observer{d}

// 	// When.
// 	observer.EXPECT().Update("parkir1", false).Return(false).Times(1)
// 	result := observer.Update("parkir1", false)

// 	// Then.
// 	assert.Equal(t, false, result)
// 	// assert.NoError(t, err)
// }

func TestAttendant_WithMock_AddCar_GivenAddingCar_ShouldReturnTicket(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parkingLot := mockparking.NewMockParkingItf(ctrl)
	att := NewAttendant("nama 1", parkingLot, "default")

	// When
	// parkingLot.EXPECT().GetStatus().Return(false).Times(1)

	// parkingLot.EXPECT().CheckCarExist(car2).Return("Car is not recognized", nil).Times(1)

	parkingLot.EXPECT().GetName().Return("parkir-1").Times(2)

	parkingLot.EXPECT().AddCar(&parkingSystem, car2).Return("ticket#5", nil).Times(1)

	result, err := att.AddCar(&parkingSystem, car2)

	// Then

	assert.Equal(t, "ticket#5", result)
	assert.NoError(t, err)
}

func TestAttendant_WithMock_AddCar_GivenAddingCarDuplicate_ShouldReturnErrorDuplicate(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parkingLot := mockparking.NewMockParkingItf(ctrl)
	att := NewAttendant("nama 1", parkingLot, "default")

	// When

	//parkingLot.EXPECT().CheckCarExist(car).Return("Car already parked", carAlreadyParked).Times(1)

	// parkingLot.EXPECT().GetName().Return("parkir-1").Times(1)

	// parkingLot.EXPECT().AddCar(&parkingSystem, car).Return("ticket#0", nil).Times(1)

	result, err := att.AddCar(&parkingSystem, car)

	// Then

	assert.Equal(t, "Car already parked", result)
	assert.Error(t, err, carAlreadyParked)
}

// func TestAttendant_WithMock_AddCar_GivenAddingWhenFull_ShouldReturnTicket(t *testing.T) {
// 	// Given.
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	att := mockparking.NewMockParkingItf(ctrl)
// 	// observer := &Observer{d}

// 	// When.
// 	att.EXPECT().AddCar(parkingSystem, car).Return("ticket#0", nil).Times(1)
// 	result, err := att.AddCar(&parkingSystem, car)

// 	// Then.
// 	assert.Equal(t, "ticket#0", result)
// 	assert.NoError(t, err)
// }

// func TestAttendant_WithMock_AddCar_GivenAdding_ShouldReturnTicket(t *testing.T) {
// 	// Given.
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	att := mockparking.NewMockParkingItf(ctrl)
// 	// observer := &Observer{d}

// 	// When.
// 	att.EXPECT().AddCar(parkingSystem, car).Return("ticket#0", nil).Times(1)
// 	result, err := att.AddCar(&parkingSystem, car)

// 	// Then.
// 	assert.Equal(t, "ticket#0", result)
// 	assert.NoError(t, err)
// }

func TestAttendant_WithMock_GetCar_GivenValidTicket_ShouldReturnCar(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parkingLot := mockparking.NewMockParkingItf(ctrl)
	att := NewAttendant("nama 1", parkingLot, "default")

	// When

	parkingLot.EXPECT().GetCar(&parkingSystem, "ticket#0").Return("Car successfully unparked", nil).Times(1)

	result, err := att.GetCar(&parkingSystem, "ticket#0")

	// Then

	assert.Equal(t, "Car successfully unparked", result)
	assert.NoError(t, err)
}

// func TestAttendant_WithMock_ArrangingParkingLot(t *testing.T) {
// 	// Given
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	parkingLot := mockparking.NewMockParkingItf(ctrl)
// 	att := NewAttendant("nama 1", parkingLot, true)
// 	parkir := NewParking("parkir 1", 1)
// 	parkir2 := NewParking("parkir 2", 2)

// 	// parkir3 := NewParking("parkir 3", 2)

// 	// When

// 	parkingLot.EXPECT().GetMaximum().Return(1).Times(2)
// 	// parkingLot.EXPECT().GetName().Return(gomock.Any()).Times(1)

// 	att.AddParkingLot(parkir, parkir2)

// 	// Then

// 	// assert.Equal(t, "Car successfully unparked", result)
// 	// assert.NoError(t, err)
// }

func TestParking_GetName(t *testing.T) {
	type fields struct {
		Name         string
		MaxLot       int
		LotCounter   int
		Status       bool
		Car          []model.Car
		Ticket       []Ticket
		observerList []observer.Observer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Coba testing ngambil nama parkir 1",
			fields: fields{
				Name: "parkir 1",
			},
			want: "parkir 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parking{
				Name:         tt.fields.Name,
				MaxLot:       tt.fields.MaxLot,
				LotCounter:   tt.fields.LotCounter,
				Status:       tt.fields.Status,
				Car:          tt.fields.Car,
				Ticket:       tt.fields.Ticket,
				observerList: tt.fields.observerList,
			}
			if got := p.GetName(); got != tt.want {
				t.Errorf("Parking.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParking_GetMaximum(t *testing.T) {
	type fields struct {
		Name         string
		MaxLot       int
		LotCounter   int
		Status       bool
		Car          []model.Car
		Ticket       []Ticket
		observerList []observer.Observer
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Coba testing ngambil max lot parkir 1 > 0",
			fields: fields{
				Name:   "parkir 1",
				MaxLot: 2,
			},
			want: 2,
		},
		{
			name: "Coba testing ngambil max lot parkir 2 == 0",
			fields: fields{
				Name:   "parkir 2",
				MaxLot: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parking{
				Name:         tt.fields.Name,
				MaxLot:       tt.fields.MaxLot,
				LotCounter:   tt.fields.LotCounter,
				Status:       tt.fields.Status,
				Car:          tt.fields.Car,
				Ticket:       tt.fields.Ticket,
				observerList: tt.fields.observerList,
			}
			if got := p.GetMaximum(); got != tt.want {
				t.Errorf("Parking.GetMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParking_GetStatus(t *testing.T) {
	type fields struct {
		Name             string
		MaxLot           int
		LotCounter       int
		FreeSpaceCounter int
		Status           bool
		Car              []model.Car
		Ticket           []Ticket
		observerList     []observer.Observer
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Coba testing ngambil status parkir 1 true",
			fields: fields{
				Name:   "parkir 1",
				MaxLot: 2,
				Status: true,
			},
			want: true,
		},
		{
			name: "Coba testing ngambil status parkir 2 false",
			fields: fields{
				Name:   "parkir 2",
				MaxLot: 2,
				Status: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parking{
				Name:             tt.fields.Name,
				MaxLot:           tt.fields.MaxLot,
				LotCounter:       tt.fields.LotCounter,
				FreeSpaceCounter: tt.fields.FreeSpaceCounter,
				Status:           tt.fields.Status,
				Car:              tt.fields.Car,
				Ticket:           tt.fields.Ticket,
				observerList:     tt.fields.observerList,
			}
			if got := p.GetStatus(); got != tt.want {
				t.Errorf("Parking.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParking_GetFreeSpace(t *testing.T) {
	type fields struct {
		Name             string
		MaxLot           int
		LotCounter       int
		FreeSpaceCounter int
		Status           bool
		Car              []model.Car
		Ticket           []Ticket
		observerList     []observer.Observer
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Coba testing ngambil free lot parkir 1 > 0",
			fields: fields{
				Name:             "parkir 1",
				FreeSpaceCounter: 2,
			},
			want: 2,
		},
		{
			name: "Coba testing ngambil free lot parkir 2 == 0",
			fields: fields{
				Name:             "parkir 2",
				FreeSpaceCounter: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parking{
				Name:             tt.fields.Name,
				MaxLot:           tt.fields.MaxLot,
				LotCounter:       tt.fields.LotCounter,
				FreeSpaceCounter: tt.fields.FreeSpaceCounter,
				Status:           tt.fields.Status,
				Car:              tt.fields.Car,
				Ticket:           tt.fields.Ticket,
				observerList:     tt.fields.observerList,
			}
			if got := p.GetFreeSpace(); got != tt.want {
				t.Errorf("Parking.GetFreeSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendant_AddParkingLot(t *testing.T) {
	type fields struct {
		id                      string
		Name                    string
		ParkingLot              []model.ParkingItf
		ParkingLotSort          []model.ParkingItf
		ParkingLotSortFreeSpace []model.ParkingItf
		Car                     *model.Car
		Ticket                  string
		ParkirFull              ParkirFull
		styleSort               string
	}
	type args struct {
		parkir []model.ParkingItf
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Coba testing ngisi 1 parking lot",
			fields: fields{
				Name: "att 1",
			},
			args: args{
				parkir: []model.ParkingItf{&Parking{Name: "parkir 1"}}},
		},
		{
			name: "Coba testing ngisi lebih dari 1 parking lot",
			fields: fields{
				Name: "att 2",
			},
			args: args{
				parkir: []model.ParkingItf{&Parking{Name: "parkir 1"}, &Parking{Name: "parkir 2"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Attendant{
				id:                      tt.fields.id,
				Name:                    tt.fields.Name,
				ParkingLot:              tt.fields.ParkingLot,
				ParkingLotSort:          tt.fields.ParkingLotSort,
				ParkingLotSortFreeSpace: tt.fields.ParkingLotSortFreeSpace,
				Car:                     tt.fields.Car,
				Ticket:                  tt.fields.Ticket,
				ParkirFull:              tt.fields.ParkirFull,
				styleSort:               tt.fields.styleSort,
			}
			a.AddParkingLot(tt.args.parkir...)
		})
	}
}

func TestAttendant_Update(t *testing.T) {
	type fields struct {
		id                      string
		Name                    string
		ParkingLot              []model.ParkingItf
		ParkingLotSort          []model.ParkingItf
		ParkingLotSortFreeSpace []model.ParkingItf
		Car                     *model.Car
		Ticket                  string
		ParkirFull              ParkirFull
		styleSort               string
	}
	type args struct {
		name   string
		status bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Coba testing update parkir 1 true",
			fields: fields{
				Name:       "att 1",
				ParkirFull: ParkirFull{},
			},
			args: args{
				name:   "parkir 1",
				status: true,
			},
			want: true,
		},
		{
			name: "Coba testing update parkir 2 false",
			fields: fields{
				Name:       "att 2",
				ParkirFull: ParkirFull{},
			},
			args: args{
				name:   "parkir 2",
				status: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Attendant{
				id:                      tt.fields.id,
				Name:                    tt.fields.Name,
				ParkingLot:              tt.fields.ParkingLot,
				ParkingLotSort:          tt.fields.ParkingLotSort,
				ParkingLotSortFreeSpace: tt.fields.ParkingLotSortFreeSpace,
				Car:                     tt.fields.Car,
				Ticket:                  tt.fields.Ticket,
				ParkirFull:              tt.fields.ParkirFull,
				styleSort:               tt.fields.styleSort,
			}
			if got := a.Update(tt.args.name, tt.args.status); got != tt.want {
				t.Errorf("Attendant.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
