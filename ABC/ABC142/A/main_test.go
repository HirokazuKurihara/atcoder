package main

import "testing"

func TestABC142A(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "ok",
			args: args{4},
			want: 0.5000000000,
		},
		{
			name: "ok",
			args: args{5},
			want: 0.6000000000,
		},
		{
			name: "ok",
			args: args{1},
			want: 1.0000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC142A(tt.args.n); got != tt.want {
				t.Errorf("ABC142A() = %v, want %v", got, tt.want)
			}
		})
	}
}
