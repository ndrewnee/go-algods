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
		{name: "1", args: args{x: 0, y: 1}, want: 1},
		{name: "2", args: args{x: 0, y: 0}, want: 0},
		{name: "3", args: args{x: 6, y: 0}, want: 2},
		{name: "4", args: args{x: 65280, y: 255}, want: 16},
		{name: "5", args: args{x: 2, y: 0}, want: 1},
		{name: "6", args: args{x: 4, y: 5}, want: 1},
		{name: "7", args: args{x: 127, y: 126}, want: 1},
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
		{name: "1", args: args{numbers: []int{1, 2, 2, 3, 3, 4, 4}}, want: 1},
		{name: "2", args: args{numbers: []int{2, 2, 1, 3, 1, 3, 4, 4, 5}}, want: 5},
		{name: "3", args: args{numbers: []int{2, 2, 3, 6, 3, 4, 4}}, want: 6},
		{name: "4", args: args{numbers: []int{12, 1, 12, 3, 1, 2, 3}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.args.numbers)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
		nums []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
			nums: []int{0, 1, 2, 3, 4},
		},
		{
			name: "2",
			args: args{
				nums: []int{-1, 0, 0, 0, 0, 3, 3},
			},
			want: 3,
			nums: []int{-1, 0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.nums, tt.args.nums[:tt.want])
		})
	}
}
