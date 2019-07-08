package roz_test

import (
	"fmt"
	"testing"

	"github.com/stevecallear/roz"
)

func ExampleNew() {
	s := roz.New("My", "URL Slug")
	fmt.Println(s)
	// Output: my-url-slug
}

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		values []string
		exp    string
	}{
		{
			name:   "should return empty for empty values",
			values: []string{},
			exp:    "",
		},
		{
			name:   "should not modify a valid slug",
			values: []string{"a1-b2-c3"},
			exp:    "a1-b2-c3",
		},
		{
			name:   "should remove duplicate dashes",
			values: []string{"--a1---b2-"},
			exp:    "a1-b2",
		},
		{
			name:   "should convert to lower case",
			values: []string{"A1 B2"},
			exp:    "a1-b2",
		},
		{
			name:   "should simplify special characters",
			values: []string{"Şpècïal Ĉhãractêrs ✓"},
			exp:    "special-characters",
		},
		{
			name:   "should remove white space",
			values: []string{" a1\nb2\t"},
			exp:    "a1-b2",
		},
		{
			name:   "should combine multiple values",
			values: []string{"  ", " a1\t", "--", "\nb2  "},
			exp:    "a1-b2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(st *testing.T) {
			act := roz.New(tt.values...)

			if act != tt.exp {
				t.Errorf("got %s, expected %s", act, tt.exp)
			}
		})
	}
}
