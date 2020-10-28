package main

import "testing"

func TestABC179B(t *testing.T) {
	type args struct {
		n     int
		slice [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yes",
			args: args{5, [][]int{{1, 2}, {6, 6}, {4, 4}, {3, 3}, {3, 2}}},
			want: "Yes",
		},
		{
			name: "No",
			args: args{5, [][]int{{1, 1}, {2, 2}, {3, 4}, {5, 5}, {6, 6}}},
			want: "No",
		},
		{
			name: "Yes",
			args: args{6, [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC179B(tt.args.n, tt.args.slice); got != tt.want {
				t.Errorf("ABC179B() = %v, want %v", got, tt.want)
			}
		})
	}
}
