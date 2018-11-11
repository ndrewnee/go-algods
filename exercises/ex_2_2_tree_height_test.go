package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeight(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "#1",
			args: args{
				tree: []int{4, -1, 4, 1, 1},
			},
			want: 3,
		},
		{
			name: "#2",
			args: args{
				tree: []int{-1, 0, 4, 0, 3},
			},
			want: 4,
		},
		{
			name: "#3",
			args: args{
				tree: []int{9, 7, 5, 5, 2, 9, 9, 9, 2, -1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Height(tt.args.tree)
			assert.Equal(t, tt.want, got)
		})
	}
}
