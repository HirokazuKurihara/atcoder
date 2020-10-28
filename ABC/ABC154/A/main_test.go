package main

import "testing"

func TestABC154A(t *testing.T) {
	type args struct {
		s string
		t string
		a int
		b int
		u string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "ok",
			args:  args{"red", "blue", 3, 4, "red"},
			want:  2,
			want1: 4,
		},
		{
			name:  "ok",
			args:  args{"red", "blue", 5, 5, "blue"},
			want:  5,
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ABC154A(tt.args.s, tt.args.t, tt.args.a, tt.args.b, tt.args.u)
			if got != tt.want {
				t.Errorf("ABC154A() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ABC154A() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
