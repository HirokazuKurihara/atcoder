package main

import (
	"reflect"
	"testing"
)

func TestABC180C(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ok",
			args: args{6},
			want: []int{1, 2, 3, 6},
		},
		{
			name: "ok",
			args: args{720},
			want: []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 30, 36, 40, 45, 48, 60, 72, 80, 90, 120, 144, 180, 240, 360, 720},
		},
		{
			name: "ok",
			args: args{1000000007},
			want: []int{1, 1000000007},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC180C(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ABC180C() = %v, want %v", got, tt.want)
			}
		})
	}
}
