package main

import (
	"reflect"
	"testing"
)

func TestABC177C(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{[]int{1, 2, 3}},
			want: 11,
		},
		{
			name: "ok",
			args: args{[]int{141421356, 17320508, 22360679, 244949}},
			want: 437235829,
		},
		{
			name: "ok",
			args: args{[]int{141421356, 17320508, 22360679, 244949, 153945}},
			want: 976696350,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC177C(tt.args.slice); got != tt.want {
				t.Errorf("ABC177C() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createCumulativeSumSlice(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ok",
			args: args{[]int{1, 2, 3}},
			want: []int{0, 1, 3, 6},
		},
		{
			name: "ok",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: []int{0, 1, 3, 6, 10, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createCumulativeSumSlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cumulativeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
