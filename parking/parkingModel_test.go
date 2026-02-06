package parking

import (
	"reflect"
	"testing"
)

func TestNewCar(t *testing.T) {
	type args struct {
		tipe     string
		colour   string
		plateNum string
	}
	tests := []struct {
		name string
		args args
		want *Car
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCar(tt.args.tipe, tt.args.colour, tt.args.plateNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendant_AddParkingLot(t *testing.T) {
	type fields struct {
		Name         string
		ParkingLot   []*Parking
		Car          *Car
		Ticket       string
		Subscription Subscription
	}
	type args struct {
		parkir *Parking
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Attendant{
				Name:         tt.fields.Name,
				ParkingLot:   tt.fields.ParkingLot,
				Car:          tt.fields.Car,
				Ticket:       tt.fields.Ticket,
				Subscription: tt.fields.Subscription,
			}
			a.AddParkingLot(tt.args.parkir)
		})
	}
}
