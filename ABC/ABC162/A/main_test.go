package main

import "testing"

func TestABC162A(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yes",
			args: args{"117"},
			want: "Yes",
		},
		{
			name: "No",
			args: args{"123"},
			want: "No",
		}, {
			name: "Yes",
			args: args{"777"},
			want: "Yes",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC162A(tt.args.n); got != tt.want {
				t.Errorf("ABC162A() = %v, want %v", got, tt.want)
			}
		})
	}
}
