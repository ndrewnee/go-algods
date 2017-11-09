package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingDistance(t *testing.T) {
	type args struct {
		x uint64
		y uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"Valid",
			args{
				151,
				180,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HammingDistance(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
		})
	}
}
