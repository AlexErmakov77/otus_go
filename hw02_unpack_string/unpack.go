package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var s2 string
	var s3 string
	var sb strings.Builder
	var t bool
	t = true
	if s == "" {
		return "", nil
	} else {
		sr := []rune(s)
		lenS := utf8.RuneCountInString(string(sr)) - 1

		for i := 0; i < lenS; i++ {
			if i == 0 && unicode.IsDigit(sr[i]) {
				t = false
			}
			if i < lenS-1 && unicode.IsDigit(sr[i]) && unicode.IsDigit(sr[i+1]) {
				t = false
			}
		}

		switch t {
		case false:
			return "", ErrInvalidString

		case true:
			for i := 0; i < lenS; i++ {
				if !unicode.IsDigit(sr[i]) && !unicode.IsDigit(sr[i+1]) {
					sb.WriteString(string(sr[i]))
				}

				if !unicode.IsDigit(sr[i]) && unicode.IsDigit(sr[i+1]) {
					counti, err := strconv.Atoi(string(sr[i+1]))
					if err != nil {
						panic(err)
					}
					s3 = (strings.Repeat(string(sr[i]), counti))
					sb.WriteString(s3)
				}

				if unicode.IsDigit(sr[i]) && !unicode.IsDigit(sr[i+1]) {
					continue
				}
			}

			if !unicode.IsDigit(sr[lenS]) {
				sb.WriteString(string(sr[lenS]))
			}

			s2 = sb.String()
		}
		return s2, nil
	}
}
