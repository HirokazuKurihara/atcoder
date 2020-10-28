package main

import "testing"

func TestABC175A(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2",
			args: args{"RRS"},
			want: 2,
		},
		{
			name: "0",
			args: args{"SSS"},
			want: 0,
		},
		{
			name: "1",
			args: args{"RSR"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC175A(tt.args.s); got != tt.want {
				t.Errorf("ABC175A() = %v, want %v", got, tt.want)
			}
		})
	}
}
