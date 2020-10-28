package main

import "testing"

func TestABC137A(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "plus",
			args: args{-13, 3},
			want: -10,
		},
		{
			name: "minus",
			args: args{1, -33},
			want: 34,
		},
		{
			name: "mul",
			args: args{13, 3},
			want: 39,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC137A(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ABC137A() = %v, want %v", got, tt.want)
			}
		})
	}
}
