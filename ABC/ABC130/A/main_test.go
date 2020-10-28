package main

import "testing"

func TestABC130A(t *testing.T) {
	type args struct {
		x int
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "less",
			args: args{3, 5},
			want: 0,
		},
		{
			name: "greater",
			args: args{7, 5},
			want: 10,
		},
		{
			name: "greater",
			args: args{6, 6},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC130A(tt.args.x, tt.args.a); got != tt.want {
				t.Errorf("ABC130A() = %v, want %v", got, tt.want)
			}
		})
	}
}
