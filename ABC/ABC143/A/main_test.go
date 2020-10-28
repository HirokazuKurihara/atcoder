package main

import "testing"

func TestABC143A(t *testing.T) {
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
			args: args{12, 4},
			want: 4,
		},
		{
			name: "ok",
			args: args{20, 15},
			want: 0,
		},
		{
			name: "ok",
			args: args{20, 30},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC143A(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ABC143A() = %v, want %v", got, tt.want)
			}
		})
	}
}
