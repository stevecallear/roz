package roz

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

var (
	chars = []*unicode.RangeTable{unicode.Letter, unicode.Number}
	seps  = []*unicode.RangeTable{unicode.White_Space, unicode.Dash}
)

// New returns a new slug for the specified strings
func New(values ...string) string {
	str := norm.NFKD.String(strings.Join(values, "-"))
	buf := make([]rune, 0, len(str))

	sep := false
	for _, r := range str {
		switch {
		case unicode.IsOneOf(chars, r):
			buf = append(buf, unicode.ToLower(r))
			sep = false
		case !sep && unicode.IsOneOf(seps, r):
			buf = append(buf, '-')
			sep = true
		}
	}

	return strings.Trim(string(buf), "-")
}
