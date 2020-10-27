package main

import "testing"

func TestABC180B(t *testing.T) {
	type args struct {
		n     int
		slice []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
		want2 float64
	}{
		{
			name:  "ok",
			args:  args{2, []float64{2, 1}},
			want:  3,
			want1: 2.236067977499790,
			want2: 2,
		},
		{
			name:  "ok",
			args:  args{10, []float64{3, -1, -4, 1, -5, 9, 2, -6, 5, -3}},
			want:  39,
			want1: 14.387494569938159,
			want2: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ABC180B(tt.args.n, tt.args.slice)
			if got != tt.want {
				t.Errorf("ABC180B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ABC180B() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ABC180B() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
