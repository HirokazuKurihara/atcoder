package main

import "testing"

func TestABC150A(t *testing.T) {
	type args struct {
		k int
		x int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yes",
			args: args{2, 900},
			want: "Yes",
		},
		{
			name: "No",
			args: args{1, 501},
			want: "No",
		},
		{
			name: "Yes",
			args: args{4, 2000},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC150A(tt.args.k, tt.args.x); got != tt.want {
				t.Errorf("ABC150A() = %v, want %v", got, tt.want)
			}
		})
	}
}
