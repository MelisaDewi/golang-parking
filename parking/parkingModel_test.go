package parking

import (
	"fmt"
	mockobserver "parkingLot/mock/mock_observer"
	"parkingLot/model"
	"parkingLot/observer"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	parkingSystem = model.NewParkingSystem()
	ps            = NewParkingSystem()
	fs            = SortFreeSpace{}
	ml            = SortMaxLot{}
	sq            = SortSequential{}

	att         = NewAttendantNoLot("att 0", sq)
	att1        = NewAttendant("att 1", parking, sq)
	att2        = NewAttendant("att 2", parking, ml)
	att3        = NewAttendant("att 3", parking, fs)
	parking, _  = NewParking("parkir 1", 2)
	parking2, _ = NewParking("parkir 2", 2)
	parking3, _ = NewParking("parkir 3", 3)
	parking4, _ = NewParking("parkir 4", -2)
	car         = model.NewCar("apa", "putih", "1111")
	car2        = model.NewCar("apa", "putih", "2222")
	car3        = NewCar("tipe 1", "warna 1", "1212")
	CarNum      = map[string]struct{}{}
	TicketNum   = map[string]struct{}{}
	// CarNum["1212"] = struct{}{}
)

func TestObserver_WithMock_Update_GivenParkir1FullButNoSub_ShouldReturnTheCorrectResult(t *testing.T) {
	// Given.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	observer := mockobserver.NewMockObserver(ctrl)
	parkir, _ := NewParking("parkir 1", 1)

	observer.EXPECT().Update("parkir 1", true).Return(true).Times(0)

	result, err := parkir.AddCar(&parkingSystem, car)

	// Then.
	assert.Equal(t, "#0", result)
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
	assert.Equal(t, "#1", result)
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

	result, err := parkir.GetCar(&parkingSystem, "#2")

	// Then.
	assert.Equal(t, "Car with plate 1111 has been unparked", result)
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

	result, err := parkir.GetCar(&parkingSystem, "#3")

	// Then.
	assert.Equal(t, "Car with plate 1111 has been unparked", result)
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
	parkir.GetCar(&parkingSystem, "#4")
	parkir.Register(observer)

	observer.EXPECT().Update("parkir 1", true).Return(true).Times(1)

	result, err := parkir.AddCar(&parkingSystem, car)

	// Then.
	assert.Equal(t, "#5", result)
	assert.NoError(t, err)
}

/*
func TestAttendant_WithMock_AddCar_GivenAddingCar_ShouldReturnTicket(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parkingLot := mockparking.NewMockParkingItf(ctrl)
	att := NewAttendant("nama 1", parkingLot)

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
*/

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
		styleSort               ParkingStyle
		ParkingStyle            ParkingTypeItf
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
			name: "COba add parking lot jumlah = 1",
			fields: fields{
				Name: "att 1",
			},
			args: args{
				parkir: []model.ParkingItf{parking},
			},
		},
		{
			name: "COba add parking lot jumlah > 1",
			fields: fields{
				Name: "att 1",
			},
			args: args{
				parkir: []model.ParkingItf{parking, parking2},
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
				ParkingStyle:            tt.fields.ParkingStyle,
			}
			a.AddParkingLot(tt.args.parkir...)
		})
	}
}

func TestParking_GetOccupiedSpace(t *testing.T) {
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
			name: "Coba testing ngambil occupied space parkir 1 > 0",
			fields: fields{
				Name:       "parkir 1",
				LotCounter: 2,
			},
			want: 2,
		},
		{
			name: "Coba testing ngambil occupied space parkir 2 == 0",
			fields: fields{
				Name:       "parkir 2",
				LotCounter: 0,
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
			if got := p.GetOccupiedSpace(); got != tt.want {
				t.Errorf("Parking.GetOccupiedSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortSequential_Priority(t *testing.T) {
	type args struct {
		a     *Attendant
		style ParkingStyle
	}
	tests := []struct {
		name string
		s    SortSequential
		args args
		want []model.ParkingItf
	}{
		{
			name: "Coba test sequential priority",
			s:    SortSequential{},
			args: args{
				a:     att1,
				style: att1.styleSort,
			},
			want: []model.ParkingItf{parking},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SortSequential{}
			if got := s.Priority(tt.args.a, tt.args.style); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortSequential.Priority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortMaxLot_Priority(t *testing.T) {
	type args struct {
		a     *Attendant
		style ParkingStyle
	}
	tests := []struct {
		name string
		s    SortMaxLot
		args args
		want []model.ParkingItf
	}{
		{
			name: "Coba test max lot priority",
			s:    SortMaxLot{},
			args: args{
				a:     att2,
				style: att2.styleSort,
			},
			want: []model.ParkingItf{parking},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SortMaxLot{}
			if got := s.Priority(tt.args.a, tt.args.style); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortMaxLot.Priority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortFreeSpace_Priority(t *testing.T) {
	att3.AddParkingLot(parking2)
	att3.AddParkingLot(parking3)
	type args struct {
		a     *Attendant
		style ParkingStyle
	}
	tests := []struct {
		name string
		s    SortFreeSpace
		args args
		want []model.ParkingItf
	}{
		{
			name: "Coba test free space priority",
			s:    SortFreeSpace{},
			args: args{
				a:     att3,
				style: att3.styleSort,
			},
			want: []model.ParkingItf{parking3, parking3, parking, parking2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SortFreeSpace{}
			if got := s.Priority(tt.args.a, tt.args.style); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortFreeSpace.Priority() = %v, want %v", got, tt.want)
			}
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
		styleSort               ParkingStyle
		ParkingStyle            ParkingTypeItf
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
			name: "Coba test update parkir 1 full",
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
			name: "Coba test update parkir 1 not full",
			fields: fields{
				Name:       "att 2",
				ParkirFull: ParkirFull{},
			},
			args: args{
				name:   "parkir 1",
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
				ParkingStyle:            tt.fields.ParkingStyle,
			}
			if got := a.Update(tt.args.name, tt.args.status); got != tt.want {
				t.Errorf("Attendant.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendant_GetID(t *testing.T) {
	type fields struct {
		id                      string
		Name                    string
		ParkingLot              []model.ParkingItf
		ParkingLotSort          []model.ParkingItf
		ParkingLotSortFreeSpace []model.ParkingItf
		Car                     *model.Car
		Ticket                  string
		ParkirFull              ParkirFull
		styleSort               ParkingStyle
		ParkingStyle            ParkingTypeItf
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Coba cari id attendant",
			fields: fields{
				id:   "1",
				Name: "att 2",
			},
			want: "1",
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
				ParkingStyle:            tt.fields.ParkingStyle,
			}
			if got := a.GetID(); got != tt.want {
				t.Errorf("Attendant.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendant_ChangeStyle(t *testing.T) {
	type fields struct {
		id                      string
		Name                    string
		ParkingLot              []model.ParkingItf
		ParkingLotSort          []model.ParkingItf
		ParkingLotSortFreeSpace []model.ParkingItf
		Car                     *model.Car
		Ticket                  string
		ParkirFull              ParkirFull
		styleSort               ParkingStyle
		ParkingStyle            ParkingTypeItf
	}
	type args struct {
		style ParkingStyle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Try to change parking style to sequential",
			fields: fields{
				id:   "1",
				Name: "att 2",
			},
			args: args{
				style: Sequential,
			},
		},
		{
			name: "Try to change parking style to highest max lot",
			fields: fields{
				id:   "1",
				Name: "att 2",
			},
			args: args{
				style: HighestMaxLot,
			},
		},
		{
			name: "Try to change parking style to highest free space",
			fields: fields{
				id:   "1",
				Name: "att 2",
			},
			args: args{
				style: HighestFreeSpace,
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
				ParkingStyle:            tt.fields.ParkingStyle,
			}
			a.ChangeStyle(tt.args.style)
		})
	}
}

func TestAttendant_ChangeParkingT(t *testing.T) {
	type fields struct {
		id                      string
		Name                    string
		ParkingLot              []model.ParkingItf
		ParkingLotSort          []model.ParkingItf
		ParkingLotSortFreeSpace []model.ParkingItf
		Car                     *model.Car
		Ticket                  string
		ParkirFull              ParkirFull
		styleSort               ParkingStyle
		ParkingStyle            ParkingTypeItf
	}
	type args struct {
		styleParking ParkingTypeItf
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []model.ParkingItf
	}{
		{
			name: "Try to change parking style to highest free space",
			fields: fields{
				id:           "1",
				ParkingLot:   []model.ParkingItf{parking},
				ParkingStyle: SortSequential{},
				styleSort:    Sequential,
			},
			args: args{
				styleParking: sq,
			},
			want: []model.ParkingItf{parking},
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
				ParkingStyle:            tt.fields.ParkingStyle,
			}
			if got := a.ChangeParkingT(tt.args.styleParking); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attendant.ChangeParkingT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendant_ChangeParkingType(t *testing.T) {
	type fields struct {
		id                      string
		Name                    string
		ParkingLot              []model.ParkingItf
		ParkingLotSort          []model.ParkingItf
		ParkingLotSortFreeSpace []model.ParkingItf
		Car                     *model.Car
		Ticket                  string
		ParkirFull              ParkirFull
		styleSort               ParkingStyle
		ParkingStyle            ParkingTypeItf
	}
	type args struct {
		styleParking ParkingTypeItf
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Try to change parking style to highest free space",
			fields: fields{
				id:   "1",
				Name: "att 2",
			},
			args: args{
				styleParking: sq,
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
				ParkingStyle:            tt.fields.ParkingStyle,
			}
			a.ChangeParkingType(tt.args.styleParking)
		})
	}
}

func TestParking_Deregister(t *testing.T) {
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
	type args struct {
		o observer.Observer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Try to deregister observer",
			fields: fields{
				Name:         "parkir 1",
				observerList: []observer.Observer{att},
			},
			args: args{
				o: att,
			},
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
			p.Deregister(tt.args.o)
		})
	}
}

func TestParking_CheckCarExist(t *testing.T) {
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
	type args struct {
		car *model.Car
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Try to check if car exist",
			fields: fields{
				Name: "parkir 1",
				Car:  []model.Car{*car},
			},
			args: args{
				car: car,
			},
			want:    "Car already parked",
			wantErr: true,
		},
		{
			name: "Try to check if car not exist",
			fields: fields{
				Name: "parkir 1",
				Car:  []model.Car{*car},
			},
			args: args{
				car: car2,
			},
			want:    "Car is not recognized",
			wantErr: false,
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
			got, err := p.CheckCarExist(tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parking.CheckCarExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parking.CheckCarExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSystem_CheckCarExist(t *testing.T) {
	CarNum["1212"] = struct{}{}
	type fields struct {
		CarNum map[string]struct{}
		Ticket map[string]struct{}
	}
	type args struct {
		car *Car
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Try to check if car exist",
			fields: fields{
				CarNum: CarNum,
			},
			args: args{
				car: car3,
			},
			want:    "Car already parked",
			wantErr: true,
		},
		{
			name: "Try to check if car not exist",
			fields: fields{
				CarNum: map[string]struct{}{},
			},
			args: args{
				car: car3,
			},
			want:    "Car is not recognized",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ParkingSystem{
				CarNum: tt.fields.CarNum,
				Ticket: tt.fields.Ticket,
			}
			got, err := ps.CheckCarExist(tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingSystem.CheckCarExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParkingSystem.CheckCarExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSystem_CheckTicketExist(t *testing.T) {
	TicketNum["#0"] = struct{}{}
	type fields struct {
		CarNum map[string]struct{}
		Ticket map[string]struct{}
	}
	type args struct {
		ticket string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Try to check if ticket exist",
			fields: fields{
				Ticket: TicketNum,
			},
			args: args{
				ticket: "#0",
			},
			want:    "Ticket exists",
			wantErr: false,
		},
		{
			name: "Try to check if ticket not exist",
			fields: fields{
				Ticket: map[string]struct{}{},
			},
			args: args{
				ticket: "#0",
			},
			want:    "Invalid ticket",
			wantErr: true,
		},
		{
			name: "Empty ticket string",
			fields: fields{
				Ticket: map[string]struct{}{},
			},
			args: args{
				ticket: "",
			},
			want:    "Invalid ticket",
			wantErr: true,
		},
		{
			name: "Ticket without # prefix",
			fields: fields{
				Ticket: map[string]struct{}{
					"0": {},
				},
			},
			args: args{
				ticket: "0",
			},
			want:    "Invalid ticket",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ParkingSystem{
				CarNum: tt.fields.CarNum,
				Ticket: tt.fields.Ticket,
			}
			got, err := ps.CheckTicketExist(tt.args.ticket)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingSystem.CheckTicketExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParkingSystem.CheckTicketExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendant_AddCar(t *testing.T) {
	parking5, _ := NewParking("parkir 5", 1)
	type fields struct {
		id                      string
		Name                    string
		ParkingLot              []model.ParkingItf
		ParkingLotSort          []model.ParkingItf
		ParkingLotSortFreeSpace []model.ParkingItf
		Car                     *model.Car
		Ticket                  string
		ParkirFull              ParkirFull
		styleSort               ParkingStyle
		ParkingStyle            ParkingTypeItf
	}
	type args struct {
		ps  *model.ParkingSystem
		car *model.Car
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{

		{
			name: "Adding car pass",
			fields: fields{
				id:                      "1",
				ParkingLot:              []model.ParkingItf{parking5},
				ParkingLotSort:          []model.ParkingItf{parking5},
				ParkingLotSortFreeSpace: []model.ParkingItf{parking5},
				ParkingStyle:            SortSequential{},
				styleSort:               Sequential,
			},
			args: args{
				ps:  &parkingSystem,
				car: &model.Car{PlateNum: "4321"},
			},
			want:    "#6",
			wantErr: false,
		},
		{
			name: "Adding car not pass",
			fields: fields{
				id:                      "1",
				ParkingLot:              []model.ParkingItf{},
				ParkingLotSort:          []model.ParkingItf{},
				ParkingLotSortFreeSpace: []model.ParkingItf{},
				ParkingStyle:            SortSequential{},
				styleSort:               Sequential,
			},
			args: args{
				ps:  &parkingSystem,
				car: car,
			},
			want:    "Car already parked",
			wantErr: true,
		},
		{
			name: "Parking lot full",
			fields: fields{
				id:                      "1",
				ParkingLot:              []model.ParkingItf{parking5},
				ParkingLotSort:          []model.ParkingItf{parking5},
				ParkingLotSortFreeSpace: []model.ParkingItf{parking5},
				ParkingStyle:            SortSequential{},
				styleSort:               Sequential,
			},
			args: args{
				ps:  &parkingSystem,
				car: &model.Car{PlateNum: "123334"},
			},
			want:    "Kayaknya parkirannya penuh",
			wantErr: true,
		},
		{
			name: "Skip full parking",
			fields: fields{
				id:                      "1",
				ParkingLot:              []model.ParkingItf{parking5},
				ParkingLotSort:          []model.ParkingItf{parking5},
				ParkingLotSortFreeSpace: []model.ParkingItf{parking5},
				ParkirFull:              ParkirFull{"parkir 5": true},
				ParkingStyle:            SortSequential{},
				styleSort:               Sequential,
			},
			args: args{
				ps:  &parkingSystem,
				car: &model.Car{PlateNum: "2345"},
			},
			want:    "Kayaknya parkirannya penuh",
			wantErr: true,
		},
		{
			name: "Multiple parking lots second succeeds",
			fields: fields{
				id:                      "1",
				ParkingLot:              []model.ParkingItf{parking2, parking5},
				ParkingLotSort:          []model.ParkingItf{parking2, parking5},
				ParkingLotSortFreeSpace: []model.ParkingItf{parking2, parking5},
				ParkingStyle:            SortSequential{},
				styleSort:               Sequential,
			},
			args: args{
				ps:  &parkingSystem,
				car: &model.Car{PlateNum: "5667"},
			},
			want:    "#7",
			wantErr: false,
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
				ParkingStyle:            tt.fields.ParkingStyle,
			}
			got, err := a.AddCar(tt.args.ps, tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("Attendant.AddCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Attendant.AddCar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendant_GetCar(t *testing.T) {
	parkingGetCar, _ := NewParking("parkir get car", 3)
	res, _ := parkingGetCar.AddCar(&parkingSystem, car)
	fmt.Println(res)
	res, _ = parkingGetCar.AddCar(&parkingSystem, car2)
	fmt.Println(res)
	res, _ = parkingGetCar.AddCar(&parkingSystem, model.NewCar("ini", "apa", "23445"))
	fmt.Println(res)
	fmt.Println(parkingSystem.AvailableTickets())
	type fields struct {
		id                      string
		Name                    string
		ParkingLot              []model.ParkingItf
		ParkingLotSort          []model.ParkingItf
		ParkingLotSortFreeSpace []model.ParkingItf
		Car                     *model.Car
		Ticket                  string
		ParkirFull              ParkirFull
		styleSort               ParkingStyle
		ParkingStyle            ParkingTypeItf
	}
	type args struct {
		ps     *model.ParkingSystem
		ticket string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Empty ticket string",
			fields: fields{
				id:         "1",
				ParkingLot: []model.ParkingItf{parkingGetCar},
			},
			args: args{
				ps:     &parkingSystem,
				ticket: "",
			},
			want:    "Invalid ticket",
			wantErr: true,
		},
		{
			name: "Ticket without # prefix",
			fields: fields{
				id:         "1",
				ParkingLot: []model.ParkingItf{parkingGetCar},
			},
			args: args{
				ps:     &parkingSystem,
				ticket: "8",
			},
			want:    "Invalid ticket",
			wantErr: true,
		},
		{
			name: "Ticket valid but car not in any lot",
			fields: fields{
				id:         "1",
				ParkingLot: []model.ParkingItf{parking2},
			},
			args: args{
				ps:     &parkingSystem,
				ticket: "#0",
			},
			want:    "Ga nemu mobilnya",
			wantErr: true,
		},
		{
			name: "Multiple parking lots, car in second lot",
			fields: fields{
				id:         "1",
				ParkingLot: []model.ParkingItf{parking2, parkingGetCar},
			},
			args: args{
				ps:     &parkingSystem,
				ticket: "#8",
			},
			want:    "Car with plate 1111 has been unparked",
			wantErr: false,
		},
		{
			name: "Ticket valid but not exist",
			fields: fields{
				id:         "1",
				ParkingLot: []model.ParkingItf{parking2, parkingGetCar},
			},
			args: args{
				ps:     &parkingSystem,
				ticket: "#12",
			},
			want:    "Invalid ticket",
			wantErr: true,
		},
		{
			name: "Getting car pass",
			fields: fields{
				id:                      "1",
				ParkingLot:              []model.ParkingItf{parkingGetCar},
				ParkingLotSort:          []model.ParkingItf{parkingGetCar},
				ParkingLotSortFreeSpace: []model.ParkingItf{parkingGetCar},
				ParkingStyle:            SortSequential{},
				styleSort:               Sequential,
			},
			args: args{
				ps:     &parkingSystem,
				ticket: "#9",
			},
			want:    "Car with plate 2222 has been unparked",
			wantErr: false,
		},
		{
			name: "Getting car not pass",
			fields: fields{
				id:                      "1",
				ParkingLot:              []model.ParkingItf{parking2},
				ParkingLotSort:          []model.ParkingItf{parking2},
				ParkingLotSortFreeSpace: []model.ParkingItf{parking2},
				ParkingStyle:            SortSequential{},
				styleSort:               Sequential,
			},
			args: args{
				ps:     &parkingSystem,
				ticket: "#10",
			},
			want:    "Ga nemu mobilnya",
			wantErr: true,
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
				ParkingStyle:            tt.fields.ParkingStyle,
			}
			got, err := a.GetCar(tt.args.ps, tt.args.ticket)
			if (err != nil) != tt.wantErr {
				t.Errorf("Attendant.GetCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Attendant.GetCar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeFromslice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	o1 := mockobserver.NewMockObserver(ctrl)
	o2 := mockobserver.NewMockObserver(ctrl)
	o3 := mockobserver.NewMockObserver(ctrl)

	o1.EXPECT().GetID().AnyTimes().Return("1")
	o2.EXPECT().GetID().AnyTimes().Return("2")
	o3.EXPECT().GetID().AnyTimes().Return("3")
	type args struct {
		observerList     []observer.Observer
		observerToRemove observer.Observer
	}
	tests := []struct {
		name string
		args args
		want []observer.Observer
	}{
		{
			name: "Remove non existing observer",
			args: args{
				observerList:     []observer.Observer{o1, o2},
				observerToRemove: o3,
			},
			want: []observer.Observer{o1, o2},
		},
		{
			name: "Remove observer",
			args: args{
				observerList:     []observer.Observer{o1},
				observerToRemove: o1,
			},
			want: []observer.Observer{},
		},
		{
			name: "Remove from empty observer",
			args: args{
				observerList:     []observer.Observer{},
				observerToRemove: o1,
			},
			want: []observer.Observer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeFromslice(tt.args.observerList, tt.args.observerToRemove); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeFromslice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSystem_AvailableTickets(t *testing.T) {
	type fields struct {
		CarNum map[string]struct{}
		Ticket map[string]struct{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Tickets available",
			fields: fields{
				Ticket: map[string]struct{}{
					"#0": {},
					"#1": {},
					"#2": {},
				},
			},
			want:    "#0, #1, #2",
			wantErr: false,
		},
		{
			name: "No tickets available",
			fields: fields{
				Ticket: map[string]struct{}{},
			},
			want:    "No ticket available",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ParkingSystem{
				CarNum: tt.fields.CarNum,
				Ticket: tt.fields.Ticket,
			}
			got, err := ps.AvailableTickets()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingSystem.AvailableTickets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParkingSystem.AvailableTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendant_CheckParkingExist(t *testing.T) {
	type fields struct {
		id                      string
		Name                    string
		ParkingLot              []model.ParkingItf
		ParkingLotSort          []model.ParkingItf
		ParkingLotSortFreeSpace []model.ParkingItf
		Car                     *model.Car
		Ticket                  string
		ParkirFull              ParkirFull
		styleSort               ParkingStyle
		ParkingStyle            ParkingTypeItf
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Check parking exist true",
			fields: fields{
				ParkingLot: []model.ParkingItf{parking},
			},
			want: true,
		},
		{
			name: "Check parking exist false",
			fields: fields{
				ParkingLot: []model.ParkingItf{},
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
				ParkingStyle:            tt.fields.ParkingStyle,
			}
			if got := a.CheckParkingExist(); got != tt.want {
				t.Errorf("Attendant.CheckParkingExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
