package parseargs_test

import (
	"testing"

	"github.com/nproc/parseargs-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParser(t *testing.T) {
	Convey("valid", t, func() {
		Convey("should understand simple words", func() {
			input := `my string of arguments`
			expected := []string{`my`, `string`, `of`, `arguments`}
			actual, err := parseargs.Parse(input)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})

		Convey("should understand single quote", func() {
			input := `my 'string of arguments'`
			expected := []string{`my`, `string of arguments`}
			actual, err := parseargs.Parse(input)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})

		Convey("should understand double quote", func() {
			input := `"my string" of arguments`
			expected := []string{`my string`, `of`, `arguments`}
			actual, err := parseargs.Parse(input)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})

		Convey("should understand escaped single quote", func() {
			input := `my 'string \'of\'' arguments`
			expected := []string{`my`, `string 'of'`, `arguments`}
			actual, err := parseargs.Parse(input)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})

		Convey("should understand escaped double quote", func() {
			input := `my "string \"of\"" arguments`
			expected := []string{`my`, `string "of"`, `arguments`}
			actual, err := parseargs.Parse(input)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})

		Convey("should understand double quotes inside single quotted string", func() {
			input := `my 'string "of" arguments'`
			expected := []string{`my`, `string "of" arguments`}
			actual, err := parseargs.Parse(input)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})

		Convey("should understand single quotes inside double quotted string", func() {
			input := `my "string 'of' arguments"`
			expected := []string{`my`, `string 'of' arguments`}
			actual, err := parseargs.Parse(input)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})

		Convey("should ignore consecutive spaces", func() {
			input := `my     string of    arguments`
			expected := []string{`my`, `string`, `of`, `arguments`}
			actual, err := parseargs.Parse(input)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})

		Convey("should accept tabs, newlines and cartridge returns as spaces", func() {
			input := "my\tstring\nof\rarguments"
			expected := []string{`my`, `string`, `of`, `arguments`}
			actual, err := parseargs.Parse(input)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})

		Convey("should read a one char word at the end of the input", func() {
			input := `my string of arguments 0`
			expected := []string{`my`, `string`, `of`, `arguments`, `0`}
			actual, err := parseargs.Parse(input)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
	})

	Convey("invalid", t, func() {
		Convey("should not allow for a quotted string to start right after a word", func() {
			input := `my"string" of arguments`
			expected := `invalid argument(s)`
			_, err := parseargs.Parse(input)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, expected)
		})

		Convey("should detect unexpected EOF", func() {
			input := `my "string of arguments`
			expected := `unexpected end of input`
			_, err := parseargs.Parse(input)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, expected)
		})

		Convey("should detect wrongly escaped quotes", func() {
			input := `my \\"string\\" of arguments`
			expected := `invalid argument(s)`
			_, err := parseargs.Parse(input)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, expected)
		})

		Convey("should not allow escaped spaces", func() {
			input := `my\ string of arguments`
			expected := `invalid syntax`
			_, err := parseargs.Parse(input)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, expected)
		})
	})
}
