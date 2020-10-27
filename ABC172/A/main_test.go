package main

import "testing"

func TestABC172A(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "14",
			args: args{2},
			want: 14,
		},
		{
			name: "1110",
			args: args{10},
			want: 1110,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC172A(tt.args.a); got != tt.want {
				t.Errorf("ABC172A() = %v, want %v", got, tt.want)
			}
		})
	}
}
