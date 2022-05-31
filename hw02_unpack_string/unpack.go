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
	var t bool = true

	if s == "" {
		return "", nil
	} else {
		sr := []rune(s)
		lenS := utf8.RuneCountInString(string(sr))

		for i := 0; i < lenS; i++ {
			if i == 0 && unicode.IsDigit(sr[i]) {
				t = false

			}
			if i < lenS-2 && unicode.IsDigit(sr[i]) && unicode.IsDigit(sr[i+1]) {
				t = false

			}

		}

		switch t {

		case false:
			return "", ErrInvalidString

		case true:
			for i := 1; i < lenS; {

				if unicode.IsDigit(sr[i]) {

					res, err := strconv.Atoi(string(sr[i])) //, err
					if err != nil {
						panic(err)
					}

					s3 = (strings.Repeat(string(sr[i-1]), res))
					sb.WriteString(s3)
					i += 2

				} else {
					sb.WriteString(string(sr[i-1]))
					i++
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
