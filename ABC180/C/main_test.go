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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC180C(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ABC180C() = %v, want %v", got, tt.want)
			}
		})
	}
}