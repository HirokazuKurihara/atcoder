package main

import "testing"

func TestABC178A(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{1},
			want: 0,
		},
		{
			name: "ok",
			args: args{0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC178A(tt.args.x); got != tt.want {
				t.Errorf("ABC178A() = %v, want %v", got, tt.want)
			}
		})
	}
}
