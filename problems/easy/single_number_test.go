package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
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
			t.Run("SingleNumber", func(t *testing.T) {
				got := SingleNumber(tt.args.numbers)
				assert.Equal(t, tt.want, got)
			})
			t.Run("SingleNumberDummy", func(t *testing.T) {
				dummy := SingleNumberDummy(tt.args.numbers)
				assert.Equal(t, tt.want, dummy)
			})
		})
	}
}
