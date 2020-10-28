package main

import "testing"

func TestABC149A(t *testing.T) {
	type args struct {
		k string
		x string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{"oder", "atc"},
			want: "atcoder",
		},
		{
			name: "ok",
			args: args{"humu", "humu"},
			want: "humuhumu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC149A(tt.args.k, tt.args.x); got != tt.want {
				t.Errorf("ABC149A() = %v, want %v", got, tt.want)
			}
		})
	}
}
