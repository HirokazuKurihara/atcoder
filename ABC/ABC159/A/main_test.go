package main

import "testing"

func TestABC159A(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{4, 3},
			want: 9,
		},
		{
			name: "ok",
			args: args{1, 1},
			want: 0,
		},
		{
			name: "ok",
			args: args{13, 3},
			want: 81,
		},
		{
			name: "ok",
			args: args{0, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC159A(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("ABC159A() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combination(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok1",
			args: args{4, 2},
			want: 6,
		},
		{
			name: "ok2",
			args: args{1, 2},
			want: 0,
		},
		{
			name: "ok3",
			args: args{1, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combination(tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("combination() = %v, want %v", got, tt.want)
			}
		})
	}
}
