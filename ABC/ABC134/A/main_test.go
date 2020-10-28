package main

import "testing"

func TestABC134A(t *testing.T) {
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
			args: args{4},
			want: 48,
		},
		{
			name: "ok",
			args: args{15},
			want: 675,
		},
		{
			name: "ok",
			args: args{80},
			want: 19200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC134A(tt.args.r); got != tt.want {
				t.Errorf("ABC134A() = %v, want %v", got, tt.want)
			}
		})
	}
}
