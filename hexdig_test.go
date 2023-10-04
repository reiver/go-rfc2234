package rfc2234_test

import (
	"testing"

	"sourcecode.social/reiver/go-rfc2234"
)

func TestIsHexDig(t *testing.T) {

	tests := []struct{
		Rune rune
		Expected bool
	}{
	}

	for r:=rune(0); r < '0'; r++ {
		tests = append(tests, struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: false,
		})
	}
	for r:='0'; r <= '9'; r++ {
		tests = append(tests, struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: true,
		})
	}
	for r:='9'+1; r < 'A'; r++ {
		tests = append(tests, struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: false,
		})
	}
	for r:='A'; r <= 'F'; r++ {
		tests = append(tests, struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: true,
		})
	}
	for r:='F'+1; r < rune(1024); r++ {
		tests = append(tests, struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: false,
		})
	}

	for testNumber, test := range tests {
		actual := rfc2234.IsHexDig(test.Rune)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for rfc2234.IsHexDig() is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("RUNE: (%U) %q", test.Rune, string(test.Rune))
			continue
		}
	}
}
