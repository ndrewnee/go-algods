package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxWindows(t *testing.T) {
	type args struct {
		numbers    []int
		windowSize int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should return nil because there are no numbers",
		},
		{
			name: "should return the same numbers because windowSize is 1",
			args: args{
				numbers:    []int{9, 2, 1, 3, 5, 21, 15, 6},
				windowSize: 1,
			},
			want: []int{9, 2, 1, 3, 5, 21, 15, 6},
		},
		{
			name: "should return one max because windowSize is greather than length",
			args: args{
				numbers:    []int{9, 2, 1, 3, 5, 21, 15, 6},
				windowSize: 20,
			},
			want: []int{21},
		},
		{
			name: "should return one max because windowSize is 0",
			args: args{
				numbers:    []int{9, 2, 1, 3, 5, 21, 15, 6},
				windowSize: 0,
			},
			want: []int{21},
		},
		{
			name: "should return valid max windows #1",
			args: args{
				numbers:    []int{9, 2, 1, 3, 5, 21, 15, 6},
				windowSize: 3,
			},
			want: []int{9, 3, 5, 21, 21, 21},
		},
		{
			name: "should return valid max windows #2",
			args: args{
				numbers:    []int{5, 1, 3, 2, 4, 6, 1, 7, 3, 2, 8},
				windowSize: 3,
			},
			want: []int{5, 3, 4, 6, 6, 7, 7, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMaxWindows(tt.args.numbers, tt.args.windowSize)
			assert.Equal(t, tt.want, got)
		})
	}
}
