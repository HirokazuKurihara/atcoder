package main

import "testing"

func TestABC176A(t *testing.T) {
	type args struct {
		n int
		x int
		t int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{20, 12, 6},
			want: 12,
		},
		{
			name: "ok",
			args: args{1000, 1, 1000},
			want: 1000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC176A(tt.args.n, tt.args.x, tt.args.t); got != tt.want {
				t.Errorf("ABC176A() = %v, want %v", got, tt.want)
			}
		})
	}
}
