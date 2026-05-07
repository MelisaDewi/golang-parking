package parking

import "testing"

func Test_io(t *testing.T) {
	type args struct {
		textInput string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nyoba input",
			args: args{
				textInput: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := io(tt.args.textInput); got != tt.want {
				t.Errorf("io() = %v, want %v", got, tt.want)
			}
		})
	}
}
