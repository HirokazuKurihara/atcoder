package main

import "testing"

func TestABC135A(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{2, 16},
			want: "9",
		},
		{
			name: "ng",
			args: args{0, 3},
			want: "IMPOSSIBLE",
		},
		{
			name: "ok",
			args: args{998244353, 99824435},
			want: "549034394",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC135A(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ABC135A() = %v, want %v", got, tt.want)
			}
		})
	}
}
