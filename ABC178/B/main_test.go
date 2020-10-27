package main

import "testing"

func TestABC178B(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{1, 2, 1, 1},
			want: 2,
		},
		{
			name: "ok",
			args: args{3, 5, -4, -2},
			want: -6,
		},
		{
			name: "ok",
			args: args{-1000000000, 0, -1000000000, 0},
			want: 1000000000000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC178B(tt.args.a, tt.args.b, tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("ABC178B() = %v, want %v", got, tt.want)
			}
		})
	}
}
