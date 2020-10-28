package main

import "testing"

func TestABC170A(t *testing.T) {
	type args struct {
		x [5]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{[5]int{0, 2, 3, 4, 5}},
			want: 1,
		},
		{
			name: "3",
			args: args{x: [5]int{1, 2, 0, 4, 5}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC170A(tt.args.x); got != tt.want {
				t.Errorf("ABC170A() = %v, want %v", got, tt.want)
			}
		})
	}
}
