package main

import (
	"testing"
)

func Test_sumOfMultiplesOf3And5(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"10",
			args{
				10,
			},
			23,
		},
		{
			"1000",
			args{
				1000,
			},
			233168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfMultiplesOf3And5(tt.args.max); got != tt.want {
				t.Errorf("sumOfMultiplesOf3And5() = %v, want %v", got, tt.want)
			}
		})
	}
}
