package main

import "testing"

func TestABC156A(t *testing.T) {
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
			name: "ok",
			args: args{2, 2919},
			want: 3719,
		},
		{
			name: "ok",
			args: args{22, 3051},
			want: 3051,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC156A(tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("ABC156A() = %v, want %v", got, tt.want)
			}
		})
	}
}
