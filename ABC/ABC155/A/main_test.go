package main

import "testing"

func TestABC155A(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yes",
			args: args{5, 7, 5},
			want: "Yes",
		},
		{
			name: "No",
			args: args{4, 4, 4},
			want: "No",
		},
		{
			name: "No",
			args: args{4, 9, 6},
			want: "No",
		},
		{
			name: "Yes",
			args: args{3, 3, 4},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC155A(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("ABC155A() = %v, want %v", got, tt.want)
			}
		})
	}
}
