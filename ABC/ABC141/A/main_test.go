package main

import "testing"

func TestABC141A(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Cloudy",
			args: args{"Sunny"},
			want: "Cloudy",
		},
		{
			name: "Rainy",
			args: args{"Cloudy"},
			want: "Rainy",
		},
		{
			name: "Sunny",
			args: args{"Rainy"},
			want: "Sunny",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC141A(tt.args.s); got != tt.want {
				t.Errorf("ABC141A() = %v, want %v", got, tt.want)
			}
		})
	}
}
