package model

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
		{
			name: "Testing making a car",
			args: args{
				tipe:     "Apa",
				colour:   "Ini",
				plateNum: "12234",
			},
			want: NewCar("Apa", "Ini", "12234"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCar(tt.args.tipe, tt.args.colour, tt.args.plateNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewParkingSystem(t *testing.T) {
	tests := []struct {
		name string
		want ParkingSystem
	}{
		{
			name: "parking system 1",
			want: NewParkingSystem(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParkingSystem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParkingSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSystem_CheckCarExist(t *testing.T) {
	car3 := NewCar("tipe 1", "warna 1", "1212")
	CarNum := map[string]struct{}{}
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
	TicketNum := map[string]struct{}{}
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
