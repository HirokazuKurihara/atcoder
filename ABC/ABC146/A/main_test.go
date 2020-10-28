package main

import "testing"

func TestABC146A(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "SAT",
			args: args{"SAT"},
			want: 1,
		},
		{
			name: "SUN",
			args: args{"SUN"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC146A(tt.args.s); got != tt.want {
				t.Errorf("ABC146A() = %v, want %v", got, tt.want)
			}
		})
	}
}
