package main

import "testing"

func TestABC179A(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "s",
			args: args{"apple"},
			want: "apples",
		},
		{
			name: "es",
			args: args{"bus"},
			want: "buses",
		},
		{
			name: "s",
			args: args{"box"},
			want: "boxs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC179A(tt.args.s); got != tt.want {
				t.Errorf("ABC179A() = %v, want %v", got, tt.want)
			}
		})
	}
}
