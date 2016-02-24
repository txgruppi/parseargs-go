package parseargs

import (
	"errors"
	"regexp"
	"strings"
)

var (
	whitespaceRegexp       = regexp.MustCompile("\\s")
	specialCharsRegexp     = regexp.MustCompile(`\s|"|'`)
	backSlashRemovalRegexp = regexp.MustCompile(`\\([\s"'\\])`)

	ErrInvalidArgument      = errors.New("invalid argument(s)")
	ErrInvalidSyntax        = errors.New("invalid syntax")
	ErrUnexpectedEndOfInput = errors.New("unexpected end of input")
)

func Parse(input string) ([]string, error) {
	input = strings.TrimSpace(input)
	runes := []rune(input)

	var reading bool
	var startChar rune
	var startIndex int = -1

	read := func(start, end int) []rune {
		reading = false
		startChar = 0
		startIndex = -1
		return runes[start:end]
	}

	result := []string{}
	for index, current := range runes {
		if reading && startChar == ' ' && isSpecial(current) && !isWhitespace(current) {
			return nil, ErrInvalidArgument
		}

		if !(reading || isSpecial(current)) {
			reading = true
			startChar = ' '
			startIndex = index

			if index == len(runes)-1 && startChar == ' ' {
				result = append(result, string(read(startIndex, len(runes))))
			}
			continue
		}

		if !reading && isSpecial(current) && !isWhitespace(current) {
			reading = true
			startChar = current
			startIndex = index
			continue
		}

		if !reading {
			continue
		}

		if startChar == ' ' && isWhitespace(current) {
			if !isValid(index, runes) {
				return nil, ErrInvalidSyntax
			}
			result = append(result, string(read(startIndex, index)))
			continue
		}

		if startChar == current && isSpecial(startChar) && isValid(index, runes) {
			result = append(result, string(read(startIndex+1, index)))
			continue
		}

		if index == len(runes)-1 && startChar == ' ' {
			result = append(result, string(read(startIndex, len(runes))))
			continue
		}
	}

	if startIndex >= 0 || startChar != 0 {
		return nil, ErrUnexpectedEndOfInput
	}

	for index, value := range result {
		result[index] = backSlashRemovalRegexp.ReplaceAllString(value, "$1")
	}

	return result, nil
}

func isWhitespace(r rune) bool {
	return whitespaceRegexp.MatchString(string(r))
}

func isSpecial(r rune) bool {
	return specialCharsRegexp.MatchString(string(r))
}

func isValid(index int, input []rune) bool {
	counter := 0

	for {
		if index-1-counter < 0 {
			break
		}

		if input[index-1-counter] == '\\' {
			counter += 1
			continue
		}

		break
	}

	return counter%2 == 0
}
