package main

import "testing"

func TestABC180A(t *testing.T) {
	type args struct {
		n int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{100, 1, 2},
			want: 101,
		},
		{
			name: "ok",
			args: args{100, 2, 1},
			want: 99,
		},
		{
			name: "ok",
			args: args{100, 1, 1},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC180A(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ABC180A() = %v, want %v", got, tt.want)
			}
		})
	}
}
