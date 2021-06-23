package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidChar     = errors.New("string contains invalid characters")
	ErrMultDigits      = errors.New("string contains multiple digits")
	ErrControl         = errors.New("incorrect usage of control symbol followed by letter")
	ErrStartsWithDigit = errors.New("string starts with a digit")
)

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
			return "", ErrStartsWithDigit
		case unicode.IsDigit(val) && foundControl:
			previousValue = val
			foundControl = !foundControl
		case unicode.IsDigit(val):
			countDigits++
			if countDigits > 1 {
				return "", ErrMultDigits
			}
			digit, _ := strconv.Atoi(string(val))
			bstring.WriteString(strings.Repeat(string(previousValue), digit))
			previousValue = 0
		case unicode.IsLetter(val) && !foundControl:
			if previousValue != 0 {
				bstring.WriteString(string(previousValue))
			}
			if foundControl {
				foundControl = !foundControl
			}
			previousValue = val
			countDigits = 0
		case unicode.IsLetter(val) && foundControl:
			return "", ErrControl
		case val == 92 && foundControl:
			previousValue = val
			foundControl = !foundControl
		case val == 92:
			if previousValue != 0 {
				bstring.WriteString(string(previousValue))
			}
			foundControl = !foundControl
		default:
			return "", ErrInvalidChar
		}
	}
	if previousValue != 0 {
		bstring.WriteString(string(previousValue))
	}
	return bstring.String(), nil
}
