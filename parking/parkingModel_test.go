package parking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type args struct {
		dividend float64
		divisor  float64
	}
	tests := []struct {
		name    string
		args    args // Given.
		want    float64
		wantErr string
	}{
		{
			name: "Division of two positive numbers should return the correct result",
			args: args{
				dividend: float64(20),
				divisor:  float64(10),
			},
			want:    float64(2),
			wantErr: "",
		},
		{
			name: "Division by zero should return an error",
			args: args{
				dividend: float64(20),
				divisor:  float64(0),
			},
			want:    float64(0),
			wantErr: "error cannot divide by zero",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When.
			got, err := AttAddCar(tt.args.dividend, tt.args.divisor)

			// Then.
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			}
			assert.Equal(t, tt.want, got)

		})
	}
}
