package util

import (
	"reflect"
	"testing"
)

func Test_getPermute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1},
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "12",
			args: args{
				nums: []int{1, 2},
			},
			want: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			name: "123",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name: "1234",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: [][]int{
				{1, 2, 3, 4},
				{1, 2, 4, 3},
				{1, 3, 2, 4},
				{1, 3, 4, 2},
				{1, 4, 2, 3},
				{1, 4, 3, 2},
				{2, 1, 3, 4},
				{2, 1, 4, 3},
				{2, 3, 1, 4},
				{2, 3, 4, 1},
				{2, 4, 1, 3},
				{2, 4, 3, 1},
				{3, 1, 2, 4},
				{3, 1, 4, 2},
				{3, 2, 1, 4},
				{3, 2, 4, 1},
				{3, 4, 1, 2},
				{3, 4, 2, 1},
				{4, 1, 2, 3},
				{4, 1, 3, 2},
				{4, 2, 1, 3},
				{4, 2, 3, 1},
				{4, 3, 1, 2},
				{4, 3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPermute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exclude(t *testing.T) {
	type args struct {
		nums []int
		i    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "23",
			args: args{
				nums: []int{1, 2, 3},
				i:    0,
			},
			want: []int{2, 3},
		},
		{
			name: "13",
			args: args{
				nums: []int{1, 2, 3},
				i:    1,
			},
			want: []int{1, 3},
		},
		{
			name: "12",
			args: args{
				nums: []int{1, 2, 3},
				i:    2,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exclude(tt.args.nums, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exclude() = %v, want %v", got, tt.want)
			}
		})
	}
}
