package main

import "testing"

func TestABC140A(t *testing.T) {
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
			args: args{2},
			want: 8,
		},
		{
			name: "ok",
			args: args{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC140A(tt.args.n); got != tt.want {
				t.Errorf("ABC140A() = %v, want %v", got, tt.want)
			}
		})
	}
}
