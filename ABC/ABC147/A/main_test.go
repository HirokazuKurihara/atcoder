package main

import "testing"

func TestABC147A(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "win",
			args: args{5, 7, 9},
			want: "win",
		},
		{
			name: "bust",
			args: args{13, 7, 2},
			want: "bust",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC147A(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("ABC147A() = %v, want %v", got, tt.want)
			}
		})
	}
}
