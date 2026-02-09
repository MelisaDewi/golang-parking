package parking

import (
	"fmt"
	mockobserver "parkingLot/mock/mock_observer"
	mockparking "parkingLot/mock/mock_parking"
	"parkingLot/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var parkingSystem = model.NewParkingSystem()

// var parking = model.NewParking("parkir 1", 2)
var car = model.NewCar("apa", "putih", "1111")

func TestObserver_WithMock_Update_GivenParkir1FullButNoSub_ShouldReturnTheCorrectResult(t *testing.T) {
	// Given.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	observer := mockobserver.NewMockObserver(ctrl)
	parkir := NewParking("parkir 1", 1)

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
	parkir := NewParking("parkir 1", 1)
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
	parkir := NewParking("parkir 1", 1)
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
	parkir := NewParking("parkir 1", 1)
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
	parkir := NewParking("parkir 1", 1)
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
	att := NewAttendant("nama 1", parkingLot, true)

	// When

	parkingLot.EXPECT().CheckCarExist(car).Return("Car is not recognized", nil).Times(1)

	parkingLot.EXPECT().GetName().Return("parkir-1").Times(1)

	parkingLot.EXPECT().AddCar(&parkingSystem, car).Return("ticket#0", nil).Times(1)

	result, err := att.AddCar(&parkingSystem, car)

	// Then

	assert.Equal(t, "ticket#0", result)
	assert.NoError(t, err)
}

func TestAttendant_WithMock_AddCar_GivenAddingCarDuplicate_ShouldReturnErrorDuplicate(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parkingLot := mockparking.NewMockParkingItf(ctrl)
	att := NewAttendant("nama 1", parkingLot, false)

	// When

	parkingLot.EXPECT().CheckCarExist(car).Return("Car already parked", carAlreadyParked).Times(1)

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
	att := NewAttendant("nama 1", parkingLot, false)

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
