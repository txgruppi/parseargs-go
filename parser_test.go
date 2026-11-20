package parseargs_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/txgruppi/parseargs-go"
)

func TestParse(t *testing.T) {
	valid := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "should understand simple words",
			input:    `my string of arguments`,
			expected: []string{`my`, `string`, `of`, `arguments`},
		},
		{
			name:     "should understand single quote",
			input:    `my 'string of arguments'`,
			expected: []string{`my`, `string of arguments`},
		},
		{
			name:     "should understand double quote",
			input:    `"my string" of arguments`,
			expected: []string{`my string`, `of`, `arguments`},
		},
		{
			name:     "should understand escaped single quote",
			input:    `my 'string \'of\'' arguments`,
			expected: []string{`my`, `string 'of'`, `arguments`},
		},
		{
			name:     "should understand escaped double quote",
			input:    `my "string \"of\"" arguments`,
			expected: []string{`my`, `string "of"`, `arguments`},
		},
		{
			name:     "should understand double quotes inside single quotted string",
			input:    `my 'string "of" arguments'`,
			expected: []string{`my`, `string "of" arguments`},
		},
		{
			name:     "should understand single quotes inside double quotted string",
			input:    `my "string 'of' arguments"`,
			expected: []string{`my`, `string 'of' arguments`},
		},
		{
			name:     "should ignore consecutive spaces",
			input:    `my     string of    arguments`,
			expected: []string{`my`, `string`, `of`, `arguments`},
		},
		{
			name:     "should accept tabs, newlines and cartridge returns as spaces",
			input:    "my\tstring\nof\rarguments",
			expected: []string{`my`, `string`, `of`, `arguments`},
		},
		{
			name:     "should read a one char word at the end of the input",
			input:    `my string of arguments 0`,
			expected: []string{`my`, `string`, `of`, `arguments`, `0`},
		},
		{
			name:     "should return an empty list for empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "should return an empty list for whitespace-only input",
			input:    "  \t\n ",
			expected: []string{},
		},
		{
			name:     "should trim leading and trailing spaces via tokenizer",
			input:    " a b ",
			expected: []string{`a`, `b`},
		},
		{
			name:     "should trim leading vertical tab and trailing nbsp",
			input:    "\vabc\u00a0",
			expected: []string{`abc`},
		},
		{
			name:     "should split on nbsp between words",
			input:    "abc\u00a0def",
			expected: []string{`abc`, `def`},
		},
		{
			name:     "should split on vertical tab between words",
			input:    "abc\vdef",
			expected: []string{`abc`, `def`},
		},
		{
			name:     "should preserve nbsp inside single quotes",
			input:    "'a\u00a0b'",
			expected: []string{"a\u00a0b"},
		},
		{
			name:     "should remove escaped nbsp inside double quotes",
			input:    "\"a\\\u00a0b\"",
			expected: []string{"a\u00a0b"},
		},
	}

	for _, tc := range valid {
		t.Run("valid/"+tc.name, func(t *testing.T) {
			actual, err := parseargs.Parse(tc.input)
			if err != nil {
				t.Fatalf("Parse(%q): unexpected error: %v", tc.input, err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("Parse(%q) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}

	invalid := []struct {
		name     string
		input    string
		expected error
	}{
		{
			name:     "should not allow for a quotted string to start right after a word",
			input:    `my"string" of arguments`,
			expected: parseargs.ErrInvalidArgument,
		},
		{
			name:     "should detect unexpected EOF",
			input:    `my "string of arguments`,
			expected: parseargs.ErrUnexpectedEndOfInput,
		},
		{
			name:     "should detect wrongly escaped quotes",
			input:    `my \\"string\\" of arguments`,
			expected: parseargs.ErrInvalidArgument,
		},
		{
			name:     "should not allow escaped spaces",
			input:    `my\ string of arguments`,
			expected: parseargs.ErrInvalidSyntax,
		},
		{
			name:     "should not allow trailing escaped space",
			input:    `abc\ `,
			expected: parseargs.ErrInvalidSyntax,
		},
		{
			name:     "should not allow trailing escaped nbsp",
			input:    "abc\\\u00a0",
			expected: parseargs.ErrInvalidSyntax,
		},
	}

	for _, tc := range invalid {
		t.Run("invalid/"+tc.name, func(t *testing.T) {
			_, err := parseargs.Parse(tc.input)
			if err == nil {
				t.Fatalf("Parse(%q): expected error, got nil", tc.input)
			}
			if !errors.Is(err, tc.expected) {
				t.Fatalf("Parse(%q): error = %v, want %v", tc.input, err, tc.expected)
			}
		})
	}
}
