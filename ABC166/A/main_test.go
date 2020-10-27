package main

import "testing"

func TestABC166A(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ARC",
			args: args{"ABC"},
			want: "ARC",
		},
		{
			name: "ABC",
			args: args{"ARC"},
			want: "ABC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC166A(tt.args.s); got != tt.want {
				t.Errorf("ABC166A() = %v, want %v", got, tt.want)
			}
		})
	}
}
