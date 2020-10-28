package main

import "testing"

func TestABC152A(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yes",
			args: args{3, 3},
			want: "Yes",
		},
		{
			name: "No",
			args: args{3, 2},
			want: "No",
		},
		{
			name: "Yes",
			args: args{1, 1},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC152A(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("ABC152A() = %v, want %v", got, tt.want)
			}
		})
	}
}
