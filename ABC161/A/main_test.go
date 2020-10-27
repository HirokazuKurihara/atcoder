package main

import "testing"

func TestABC161A(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{
			name:  "ok",
			args:  args{1, 2, 3},
			want:  3,
			want1: 1,
			want2: 2,
		},
		{
			name:  "ok",
			args:  args{100, 100, 100},
			want:  100,
			want1: 100,
			want2: 100,
		},
		{
			name:  "ok",
			args:  args{41, 59, 31},
			want:  31,
			want1: 41,
			want2: 59,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ABC161A(tt.args.a, tt.args.b, tt.args.c)
			if got != tt.want {
				t.Errorf("ABC161A() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ABC161A() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ABC161A() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
