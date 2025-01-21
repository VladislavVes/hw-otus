package hw02unpackstring

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "🙃0", expected: ""},
		{input: "aaф0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		// additional tests
		{input: "১২৩", expected: "১২৩"},
		{input: "১2২৩0", expected: "১১২"},
		{input: "੩4", expected: "੩੩੩੩"},
		//{input: `\\32`, expected: `\\\2`}, // 32 - number, not a digit
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "aaa\\a"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestWriteLastRune(t *testing.T) {
	tests := []struct {
		name         string
		inputCurrent rune
		inputLast    rune
		expectedR    rune
		expectedS    string
	}{
		{name: "MultiplyRune", inputCurrent: '5', inputLast: 'a', expectedR: 0, expectedS: "aaaaa"},
		{name: "ZeroRune", inputCurrent: '0', inputLast: 'a', expectedR: 0, expectedS: ""},
		{name: "WriteRune", inputCurrent: 'a', inputLast: 'b', expectedR: 'a', expectedS: "b"},
		{name: "NilCurrentRune", inputCurrent: 0, inputLast: 'a', expectedR: 0, expectedS: "a"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			builder := strings.Builder{}
			result := WriteLastRune(&builder, tc.inputCurrent, tc.inputLast)

			require.Equal(t, tc.expectedR, result)
			require.Equal(t, tc.expectedS, builder.String())
		})
	}
}

func TestCheckRune(t *testing.T) {
	tests := []struct {
		name     string
		input    rune
		expected bool
	}{
		{name: "Digit", input: '5', expected: true},
		{name: "Nil", input: 0, expected: true},
		{name: "Lane", input: '\\', expected: false},
		{name: "Symbol", input: 'a', expected: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := DigitOrNil(tc.input)

			require.Equal(t, tc.expected, result)
		})
	}
}
