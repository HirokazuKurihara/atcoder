package main

import "testing"

func TestABC160A(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
	}{
		{
			name:    "Yes",
			args:    args{"sippuu"},
			wantRet: "Yes",
		},
		{
			name:    "No",
			args:    args{"iphone"},
			wantRet: "No",
		},
		{
			name:    "Yes",
			args:    args{"coffee"},
			wantRet: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := ABC160A(tt.args.s); gotRet != tt.wantRet {
				t.Errorf("ABC160A() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
