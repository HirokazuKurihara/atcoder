package main

import "testing"

func TestABC138A(t *testing.T) {
	type args struct {
		a int
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pink",
			args: args{3200, "pink"},
			want: "pink",
		},
		{
			name: "red",
			args: args{3199, "pink"},
			want: "red",
		},
		{
			name: "red",
			args: args{4049, "red"},
			want: "red",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC138A(tt.args.a, tt.args.s); got != tt.want {
				t.Errorf("ABC138A() = %v, want %v", got, tt.want)
			}
		})
	}
}
