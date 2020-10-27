package main

import "testing"

func TestABC163A(t *testing.T) {
	type args struct {
		r int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "ok",
			args: args{1},
			want: 6.28318530717958623200,
		},
		{
			name: "ok",
			args: args{73},
			want: 458.67252742410977361942,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC163A(tt.args.r); got != tt.want {
				t.Errorf("ABC163A() = %v, want %v", got, tt.want)
			}
		})
	}
}
