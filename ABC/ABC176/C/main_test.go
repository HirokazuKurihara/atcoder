package main

import "testing"

func TestABC176C(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{[]int{2, 1, 5, 4, 3}},
			want: 4,
		},
		{
			name: "ok",
			args: args{[]int{3, 3, 3, 3, 3}},
			want: 0,
		},
		{
			name: "ok",
			args: args{[]int{3, 3, 3, 3, 3}},
			want: 0,
		},
		{
			name: "ok",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: 0,
		},
		{
			name: "ok",
			args: args{[]int{5, 4, 3, 2, 1}},
			want: 10,
		},
		{
			name: "ok",
			args: args{[]int{1, 5, 3, 2, 4}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC176C(tt.args.nums); got != tt.want {
				t.Errorf("ABC176C() = %v, want %v", got, tt.want)
			}
		})
	}
}
