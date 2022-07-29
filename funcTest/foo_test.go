package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

func Test_Do(t *testing.T) {
	tests := map[string]struct {
		letter  string
		num     int
		boolean bool
		exp     string
		expErr  error
	}{
		"valid s, valid i, true": {
			letter: "a",
			num: 2,
			boolean: true,
			exp: "[2A]",
			expErr: nil,

		},
		"valid s, valid i, false": {
			letter: "c",
			num: 1,
			boolean: false,
			exp: "[1c]",
			expErr: nil,

		},
		"valid s, skip i, true": {
			letter: "b",
			num: 21,
			boolean: true,
			exp: "B",
			expErr: nil,

		},
		"valid s, skip i, false": {
			letter: "d",
			num: 13,
			boolean: false,
			exp: "d",
			expErr: nil,

		},
		"valid s, invalid i, true": {
			letter: "a",
			num: 7,
			boolean: true,
			exp: "",
			expErr: errors.New("invalid i"),

		},
		"invalid s, valid i, false": {
			letter: "aa",
			num: 1,
			boolean: false,
			exp: "",
			expErr: errors.New("invalid s"),

		},
	}


	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
		actual, err := Do(testCase.letter, testCase.num, testCase.boolean)
		assert.Equal(t, testCase.exp, actual)
		if testCase.expErr != nil {
			assert.EqualError(t, err, testCase.expErr.Error())
		} else {
			assert.NoError(t, err)
		}
		})
	}
}