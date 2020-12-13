package study

import (
	"testing"
)

func Test_fibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "5",
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name: "7",
			args: args{
				n: 7,
			},
			want: 13,
		},
		{
			name: "40",
			args: args{
				n: 40,
			},
			want: 102334155,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args.n); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibonacciMemo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "5",
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name: "7",
			args: args{
				n: 7,
			},
			want: 13,
		},
		{
			name: "40",
			args: args{
				n: 40,
			},
			want: 102334155,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciMemo(tt.args.n); got != tt.want {
				t.Errorf("fibonacciMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.args.n); got != tt.want {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{
				m: 10,
				n: 5,
			},
			want: 5,
		},
		{
			name: "21",
			args: args{
				m: 1071,
				n: 1029,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hanoi(t *testing.T) {
	type args struct {
		n    int
		from string
		work string
		dest string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				n:    4,
				from: "A",
				work: "B",
				dest: "C",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hanoi(tt.args.n, tt.args.from, tt.args.work, tt.args.dest)
		})
	}
}
