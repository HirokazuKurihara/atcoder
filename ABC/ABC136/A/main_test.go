package main

import "testing"

func TestABC136A(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{6, 4, 3},
			want: 1,
		},
		{
			name: "ok",
			args: args{8, 3, 9},
			want: 4,
		},
		{
			name: "ok",
			args: args{12, 3, 7},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC136A(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("ABC136A() = %v, want %v", got, tt.want)
			}
		})
	}
}
