package main

import "testing"

func TestABC133A(t *testing.T) {
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
			name: "train",
			args: args{4, 2, 9},
			want: 8,
		},
		{
			name: "tax",
			args: args{4, 2, 7},
			want: 7,
		},
		{
			name: "either",
			args: args{4, 2, 8},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC133A(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ABC133A() = %v, want %v", got, tt.want)
			}
		})
	}
}
