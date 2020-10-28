package main

import "testing"

func TestABC165A(t *testing.T) {
	type args struct {
		k int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{7, 500, 600},
			want: "OK",
		},
		{
			name: "NG",
			args: args{4, 5, 7},
			want: "NG",
		},
		{
			name: "OK",
			args: args{1, 11, 11},
			want: "OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC165A(tt.args.k, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ABC165A() = %v, want %v", got, tt.want)
			}
		})
	}
}
