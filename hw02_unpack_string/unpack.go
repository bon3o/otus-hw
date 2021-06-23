package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var bstring strings.Builder
	var countDigits int
	var previousValue rune
	var foundControl bool
	if len(str) == 0 {
		return "", nil
	}
	for i, val := range str {
		switch {
		case unicode.IsDigit(val) && i == 0:
			return "", ErrInvalidString
		case unicode.IsDigit(val) && foundControl:
			previousValue = val
			foundControl = !foundControl
		case unicode.IsDigit(val):
			countDigits++
			if countDigits > 1 {
				return "", ErrInvalidString
			}
			digit, _ := strconv.Atoi(string(val))
			bstring.WriteString(strings.Repeat(string(previousValue), digit))
			previousValue = 0
		case unicode.IsLetter(val):
			if previousValue != 0 {
				bstring.WriteString(string(previousValue))
			}
			if foundControl {
				foundControl = !foundControl
			}
			previousValue = val
			countDigits = 0
		case unicode.IsLetter(val) && foundControl:
			return "", ErrInvalidString
		case val == 92 && foundControl:
			previousValue = val
			foundControl = !foundControl
		case val == 92:
			bstring.WriteString(string(previousValue))
			foundControl = !foundControl
		}
	}
	if previousValue != 0 {
		bstring.WriteString(string(previousValue))
	}
	return bstring.String(), nil
}
