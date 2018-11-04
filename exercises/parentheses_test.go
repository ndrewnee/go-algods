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
		name        string
		args        args
		parentheses Parentheses
		valid       bool
	}{
		{
			name: "should return false because str is invalid #1",
			args: args{
				str: "{][}",
			},
			parentheses: Parentheses{Rune: ']', Index: 1},
			valid:       false,
		},
		{
			name: "should return false because str is invalid #2",
			args: args{
				str: "([{}]]()",
			},
			parentheses: Parentheses{Rune: ']', Index: 5},
			valid:       false,
		},
		{
			name: "should return false because str is invalid #3",
			args: args{
				str: "{}([]",
			},
			parentheses: Parentheses{Rune: '(', Index: 2},
			valid:       false,
		},
		{
			name: "should return false because str is invalid #4",
			args: args{
				str: "][",
			},
			parentheses: Parentheses{Rune: ']', Index: 0},
			valid:       false,
		},

		{
			name: "should return true because str is valid #1",
			args: args{
				str: "{()[]([])}",
			},
			valid: true,
		},
		{
			name: "should return true because str is valid #2",
			args: args{
				str: "((({[]()[()]})))[]",
			},
			valid: true,
		},
		{
			name: "should return true because str is valid #3",
			args: args{
				str: "((({[]()[(text)]})))[test]",
			},
			valid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parentheses, valid := IsValidParentheses(tt.args.str)
			assert.Equal(t, tt.valid, valid)
			assert.Equal(t, tt.parentheses, parentheses)
		})
	}
}
