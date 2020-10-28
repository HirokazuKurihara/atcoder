package main

import "testing"

func TestABC153A(t *testing.T) {
	type args struct {
		h int
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{10, 4},
			want: 3,
		},
		{
			name: "ok",
			args: args{1, 10000},
			want: 1,
		},
		{
			name: "ok",
			args: args{10000, 1},
			want: 10000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC153A(tt.args.h, tt.args.a); got != tt.want {
				t.Errorf("ABC153A() = %v, want %v", got, tt.want)
			}
		})
	}
}
