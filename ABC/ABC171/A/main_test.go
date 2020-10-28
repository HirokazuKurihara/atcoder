package main

import "testing"

func TestABC171A(t *testing.T) {
	type args struct {
		a rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Upper",
			args: args{'B'},
			want: "A",
		},
		{
			name: "Lower",
			args: args{'a'},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC171A(tt.args.a); got != tt.want {
				t.Errorf("ABC171A() = %v, want %v", got, tt.want)
			}
		})
	}
}
