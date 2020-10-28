package main

import "testing"

func TestABC144A(t *testing.T) {
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
			args: args{2, 5},
			want: 10,
		},
		{
			name: "ng",
			args: args{5, 10},
			want: -1,
		},
		{
			name: "ok",
			args: args{9, 9},
			want: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC144A(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ABC144A() = %v, want %v", got, tt.want)
			}
		})
	}
}
