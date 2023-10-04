package rfc2234_test

import (
	"testing"

	"sourcecode.social/reiver/go-rfc2234"
)

func TestIsSubDelim(t *testing.T) {

	tests := []struct{
		Rune rune
		Expected bool
	}{
	}

	for r:=rune(0); r < rune(8192); r++ {
		test := struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: false,
		}

		if '!' == r ||
		   '$' == r ||
		   '&' == r ||
		  '\'' == r ||
		   '(' == r ||
		   ')' == r ||
		   '*' == r ||
		   '+' == r ||
		   ',' == r ||
		   ';' == r ||
		   '=' == r {
			test.Expected = true
		}

		tests = append(tests, test)
	}

	for testNumber, test := range tests {
		actual := rfc2234.IsSubDelim(test.Rune)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for rfc2234.IsSubDelim() is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("RUNE: (%U) %q", test.Rune, string(test.Rune))
			continue
		}
	}
}
