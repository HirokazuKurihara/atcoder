package main

import "testing"

func TestABC151A(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "ok",
			args: args{'a'},
			want: 'b',
		},
		{
			name: "ok",
			args: args{'y'},
			want: 'z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC151A(tt.args.c); got != tt.want {
				t.Errorf("ABC151A() = %v, want %v", got, tt.want)
			}
		})
	}
}
