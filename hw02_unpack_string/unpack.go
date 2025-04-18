package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

const saveRune rune = '\\'

func Unpack(str string) (string, error) {
	builder := strings.Builder{}
	reader := strings.NewReader(str)

	var lastR rune
	var isSaved bool

	for reader.Len() > 0 {
		r, _, err := reader.ReadRune()
		if err != nil {
			return "", err
		}

		if DigitOrNil(r) && DigitOrNil(lastR) && !isSaved {
			return "", ErrInvalidString
		}

		if lastR == saveRune && !isSaved {
			if !('0' <= r && r <= '9') && r != saveRune {
				return "", ErrInvalidString
			}

			lastR = r
			isSaved = true
		} else {
			lastR = WriteLastRune(&builder, r, lastR)
			isSaved = false
		}
	}

	if lastR != 0 {
		builder.WriteRune(lastR)
	}

	return builder.String(), nil
}

func WriteLastRune(builder *strings.Builder, r rune, lastR rune) rune {
	i, err := strconv.Atoi(string(r))

	if err == nil {
		str := strings.Repeat(string(lastR), i)
		builder.WriteString(str)
		return 0
	}

	if lastR != 0 {
		builder.WriteRune(lastR)
	}

	return r
}

func DigitOrNil(r rune) bool {
	return ('0' <= r && r <= '9') || r == 0
}
