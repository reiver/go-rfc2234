package rfc2234

// IsDigit returns true if value of 'r' matches 'ALPHA' as defined in IETF RFC-2234:
//
//	DIGIT =  %30-%39   ; 0-9
func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
