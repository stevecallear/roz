package roz

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

var (
	characters = []*unicode.RangeTable{unicode.Letter, unicode.Number}
	separators = []*unicode.RangeTable{unicode.White_Space, unicode.Dash}
)

// New returns a new slug for the specified strings
func New(vals ...string) string {
	str := norm.NFKD.String(strings.Join(vals, "-"))
	buf := make([]rune, 0, len(str))
	sep := false
	for _, r := range str {
		switch {
		case unicode.IsOneOf(characters, r):
			buf = append(buf, unicode.ToLower(r))
			sep = false
		case !sep && unicode.IsOneOf(separators, r):
			buf = append(buf, '-')
			sep = true
		}
	}
	return strings.Trim(string(buf), "-")
}
