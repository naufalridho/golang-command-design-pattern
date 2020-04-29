package main

import (
	"reflect"
	"testing"
)

func TestAlgorithm_fib(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case_1",
			args{N: 4},
			[]int{0, 1, 1 ,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Algorithm{}
			if got := a.fib(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_multiply(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case_1",
			args{x: 1, y: 2},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Algorithm{}
			if got := a.multiply(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_prime(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case_1",
			args{N: 4},
			[]int{2, 3, 5 ,7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Algorithm{}
			if got := a.prime(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_sum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case_1",
			args{x: 1, y: 2},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Algorithm{}
			if got := a.sum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}