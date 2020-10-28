package main

import "testing"

func TestABC131A(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Bad",
			args: args{"3776"},
			want: "Bad",
		},
		{
			name: "Good",
			args: args{"8080"},
			want: "Good",
		},
		{
			name: "Bad",
			args: args{"1333"},
			want: "Bad",
		},
		{
			name: "Bad",
			args: args{"0024"},
			want: "Bad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC131A(tt.args.s); got != tt.want {
				t.Errorf("ABC131A() = %v, want %v", got, tt.want)
			}
		})
	}
}
