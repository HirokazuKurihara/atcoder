package main

import "testing"

func TestABC167A(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yes",
			args: args{"chokudai", "chokudaiz"},
			want: "Yes",
		},
		{
			name: "No",
			args: args{"snuke", "snekee"},
			want: "No",
		},
		{
			name: "Yes",
			args: args{"a", "aa"},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC167A(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("ABC167A() = %v, want %v", got, tt.want)
			}
		})
	}
}
