package main

import "testing"

func TestABC177A(t *testing.T) {
	type args struct {
		d int
		t int
		s int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yes",
			args: args{1000, 15, 80},
			want: "Yes",
		},
		{
			name: "Yes",
			args: args{2000, 20, 100},
			want: "Yes",
		},
		{
			name: "No",
			args: args{10000, 1, 1},
			want: "No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC177A(tt.args.d, tt.args.t, tt.args.s); got != tt.want {
				t.Errorf("ABC177A() = %v, want %v", got, tt.want)
			}
		})
	}
}
