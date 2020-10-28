package main

import "testing"

func TestABC157A(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{5},
			want: 3,
		},
		{
			name: "ok",
			args: args{2},
			want: 1,
		},
		{
			name: "ok",
			args: args{100},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC157A(tt.args.n); got != tt.want {
				t.Errorf("ABC157A() = %v, want %v", got, tt.want)
			}
		})
	}
}
