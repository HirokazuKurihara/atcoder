package main

import "testing"

func TestABC132A(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yes",
			args: args{"ASSA"},
			want: "Yes",
		},
		{
			name: "No",
			args: args{"STOP"},
			want: "No",
		},
		{
			name: "Yes",
			args: args{"FFEE"},
			want: "Yes",
		},
		{
			name: "No",
			args: args{"FREE"},
			want: "No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC132A(tt.args.s); got != tt.want {
				t.Errorf("ABC132A() = %v, want %v", got, tt.want)
			}
		})
	}
}
