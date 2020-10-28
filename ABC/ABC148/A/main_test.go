package main

import "testing"

func TestABC148A(t *testing.T) {
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
			name: "ok",
			args: args{3, 1},
			want: 2,
		},
		{
			name: "ok",
			args: args{1, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC148A(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ABC148A() = %v, want %v", got, tt.want)
			}
		})
	}
}
