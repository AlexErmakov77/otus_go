package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	// Place your code here.

	//s := "a4bc2d5ehg0hhh0"
	//s := "арпа3пр6"
	sr := []rune(s)
	var s2 string
	var s3 string
	var sb strings.Builder
	var t bool = true

	for i := 0; i < len(sr); i++ {
		if unicode.IsDigit(sr[0]) {
			t = false

		}
		if i < len(sr)-1 && unicode.IsDigit(sr[i]) && unicode.IsDigit(sr[i+1]) {
			t = false

		}

	}

	switch t {

	case false:
		//println("error")
		return "", ErrInvalidString

	case true:
		{
			for i := 0; i < len(sr); {

				if unicode.IsDigit(sr[i+1]) {
					//s2 = string(sr[i+1])
					res, err := strconv.Atoi(string(sr[i+1])) //(s2)
					if err != nil {
						panic(err)
					}
					//println(sr[i], "*", string(sr[i]), res)

					s3 = (strings.Repeat(string(sr[i]), res))
					sb.WriteString(s3)
					i = i + 2

				} else {
					sb.WriteString(string(sr[i]))
					i = i + 1
				}
			}

			//println(sb.String())
			s2 = sb.String()
		}

	}
	return s2, nil
	//return "", nil
}
