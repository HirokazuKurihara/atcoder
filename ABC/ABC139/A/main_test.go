package main

import "testing"

func TestABC139A(t *testing.T) {
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
			args: args{"CSS", "CSR"},
			want: 2,
		},
		{
			name: "ok",
			args: args{"SSR", "SSR"},
			want: 3,
		},
		{
			name: "ok",
			args: args{"RRR", "SSS"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC139A(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("ABC139A() = %v, want %v", got, tt.want)
			}
		})
	}
}
