package main

import "testing"

func TestABC168A(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pon",
			args: args{16},
			want: "pon",
		},
		{
			name: "hon",
			args: args{2},
			want: "hon",
		},
		{
			name: "bon",
			args: args{183},
			want: "bon",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC168A(tt.args.n); got != tt.want {
				t.Errorf("ABC168A() = %v, want %v", got, tt.want)
			}
		})
	}
}
