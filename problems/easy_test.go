package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingDistance(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{0, 1}, 1},
		{"2", args{0, 0}, 0},
		{"3", args{6, 0}, 2},
		{"4", args{65280, 255}, 16},
		{"5", args{2, 0}, 1},
		{"6", args{4, 5}, 1},
		{"7", args{127, 126}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hammingDistance(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_singleNumber(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 2, 3, 3, 4, 4}}, 1},
		{"2", args{[]int{2, 2, 1, 3, 1, 3, 4, 4, 5}}, 5},
		{"3", args{[]int{2, 2, 3, 6, 3, 4, 4}}, 6},
		{"4", args{[]int{12, 1, 12, 3, 1, 2, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.args.numbers)
			assert.Equal(t, tt.want, got)
		})
	}
}
