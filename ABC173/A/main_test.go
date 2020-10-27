package main

import "testing"

func TestABC173A(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "100",
			args: args{1900},
			want: 100,
		},
		{
			name: "0",
			args: args{3000},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC173A(tt.args.n); got != tt.want {
				t.Errorf("ABC173A() = %v, want %v", got, tt.want)
			}
		})
	}
}
