package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidParentheses(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return false because str is invalid #1",
			args: args{
				str: "{][}",
			},
			want: false,
		},
		{
			name: "should return false because str is invalid #2",
			args: args{
				str: "([{}]]()",
			},
			want: false,
		},
		{
			name: "should return false because str is invalid #3",
			args: args{
				str: "{}([]",
			},
			want: false,
		},
		{
			name: "should return false because str is invalid #4",
			args: args{
				str: "][",
			},
			want: false,
		},

		{
			name: "should return true because str is valid #1",
			args: args{
				str: "{()[]([])}",
			},
			want: true,
		},
		{
			name: "should return true because str is valid #2",
			args: args{
				str: "((({[]()[()]})))[]",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidParentheses(tt.args.str)
			assert.Equal(t, tt.want, got)
		})
	}
}
