package main

import "testing"

func TestABC176B(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yes",
			args: args{"123456789"},
			want: "Yes",
		},
		{
			name: "Yes",
			args: args{"0"},
			want: "Yes",
		},
		{
			name: "No",
			args: args{"31415926535897932384626433832795028841971693993751058209749445923078164062862089986280"},
			want: "No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC176B(tt.args.n); got != tt.want {
				t.Errorf("ABC176B() = %v, want %v", got, tt.want)
			}
		})
	}
}
