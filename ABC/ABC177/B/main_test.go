package main

import "testing"

func TestABC177B(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{"cabacc", "abc"},
			want: 1,
		},
		{
			name: "ok",
			args: args{"codeforces", "atcoder"},
			want: 6,
		},
		{
			name: "ok",
			args: args{"codeatcode", "atcoder"},
			want: 7,
		},
		{
			name: "ok",
			args: args{"codatcoder", "atcoder"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC177B(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("ABC177B() = %v, want %v", got, tt.want)
			}
		})
	}
}
