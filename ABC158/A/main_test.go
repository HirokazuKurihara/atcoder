package main

import "testing"

func TestABC158A(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yes",
			args: args{"ABA"},
			want: "Yes",
		},
		{
			name: "Yes",
			args: args{"BBA"},
			want: "Yes",
		},
		{
			name: "No",
			args: args{"BBB"},
			want: "No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC158A(tt.args.s); got != tt.want {
				t.Errorf("ABC158A() = %v, want %v", got, tt.want)
			}
		})
	}
}
