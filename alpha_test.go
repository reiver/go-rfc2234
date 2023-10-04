package rfc2234_test

import (
	"testing"

	"sourcecode.social/reiver/go-rfc2234"
)

func TestIsAlpha(t *testing.T) {

	tests := []struct{
		Rune rune
		Expected bool
	}{
	}

	for r:=rune(0); r < 'A'; r++ {
		tests = append(tests, struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: false,
		})
	}
	for r:='A'; r <= 'Z'; r++ {
		tests = append(tests, struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: true,
		})
	}
	for r:='Z'+1; r < 'a'; r++ {
		tests = append(tests, struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: false,
		})
	}
	for r:='a'; r <= 'z'; r++ {
		tests = append(tests, struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: true,
		})
	}
	for r:='z'+1; r <= rune(127); r++ {
		tests = append(tests, struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: false,
		})
	}

	for testNumber, test := range tests {
		actual := rfc2234.IsAlpha(test.Rune)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for rfc2234.IsAlpha() is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("RUNE: (%U) %q", test.Rune, string(test.Rune))
			continue
		}
	}
}
