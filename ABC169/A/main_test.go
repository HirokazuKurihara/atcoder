package main

import "testing"

func TestABC169A(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10",
			args: args{2, 5},
			want: 10,
		},
		{
			name: "10000",
			args: args{100, 100},
			want: 10000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC169A(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ABC169A() = %v, want %v", got, tt.want)
			}
		})
	}
}
