package main

import "testing"

func TestABC164A(t *testing.T) {
	type args struct {
		s int
		w int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "unsafe",
			args: args{4, 5},
			want: "unsafe",
		},
		{
			name: "safe",
			args: args{100, 2},
			want: "safe",
		},
		{
			name: "unsafe",
			args: args{10, 10},
			want: "unsafe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC164A(tt.args.s, tt.args.w); got != tt.want {
				t.Errorf("ABC164A() = %v, want %v", got, tt.want)
			}
		})
	}
}
