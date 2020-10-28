package main

import "testing"

func TestABC145A(t *testing.T) {
	type args struct {
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{2},
			want: 4,
		},
		{
			name: "ok",
			args: args{100},
			want: 10000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC145A(tt.args.r); got != tt.want {
				t.Errorf("ABC145A() = %v, want %v", got, tt.want)
			}
		})
	}
}
