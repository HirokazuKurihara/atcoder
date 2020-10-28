package main

import "testing"

func TestABC174A(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No",
			args: args{25},
			want: "No",
		},
		{
			name: "Yes",
			args: args{30},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC174A(tt.args.x); got != tt.want {
				t.Errorf("ABC174A() = %v, want %v", got, tt.want)
			}
		})
	}
}
