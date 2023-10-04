package rfc2234

// IsAlpha returns true if the value of 'r' matches 'ALPHA' as defined in IETF RFC-2234:
//
//	ALPHA =  %x41-5A / %x61-7A   ; A-Z / a-z
func IsAlpha(r rune) bool {
	if 'z' < r {
		return false
	}
	if r < 'A' {
		return false
	}
	if 'Z' < r && r < 'a'{
		return false
	}

	return true
}
