package main

import "testing"

func TestABC179C(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{3},
			want: 3,
		},
		{
			name: "ok",
			args: args{100},
			want: 473,
		},
		{
			name: "ok",
			args: args{1000000},
			want: 13969985,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC179C(tt.args.n); got != tt.want {
				t.Errorf("ABC179C() = %v, want %v", got, tt.want)
			}
		})
	}
}
