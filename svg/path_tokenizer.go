package svg

import (
	"strconv"
	"unicode"
)

type tokenType int

const (
	tokenCommand tokenType = iota
	tokenNumber
)

type token struct {
	typ tokenType
	val string
}

func tokenizePath(d string) ([]token, error) {
	var tokens []token

	i := 0

	for i < len(d) {
		c := rune(d[i])

		if unicode.IsLetter(c) {
			tokens = append(tokens, token{
				typ: tokenCommand,
				val: string(c),
			})
			i++
			continue
		}

		if c == ' ' || c == ',' {
			i++
			continue
		}

		start := i
		if c == '-' || c == '+' {
			i++
		}

		for i < len(d) && (unicode.IsDigit(rune(d[i])) || d[i] == '.') {
			i++
		}

		tokens = append(tokens, token{
			typ: tokenNumber,
			val: d[start:i],
		})
	}

	return tokens, nil
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
