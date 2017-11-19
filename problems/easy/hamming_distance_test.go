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
			t.Run("HammingDistance", func(t *testing.T) {
				got := HammingDistance(tt.args.x, tt.args.y)
				assert.Equal(t, tt.want, got)
			})
			t.Run("HammingDistanceDummy", func(t *testing.T) {
				dummy := HammingDistanceDummy(tt.args.x, tt.args.y)
				assert.Equal(t, tt.want, dummy)
			})
		})
	}
}

func Test_completeStrings(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			"1st > 2nd",
			args{
				"100101101",
				"1011",
			},
			"100101101",
			"000001011",
		},
		{
			"1st < 2nd",
			args{
				"1111",
				"101101101",
			},
			"000001111",
			"101101101",
		},
		{
			"1st == 2nd",
			args{
				"111111111",
				"101101101",
			},
			"111111111",
			"101101101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := completeStrings(tt.args.s1, tt.args.s2)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
