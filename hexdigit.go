package rfc2234

// IsHexDigit returns true if value of 'r' matches 'ALPHA' as defined in IETF RFC-2234:
//
//	HEXDIGIT =  %30-%39 / %41-%46 / %61-%66    ; 0-9 / A-F / a-f
func IsHexDigit(r rune) bool {
	return ('0' <= r && r <= '9') || ('A' <= r && r <= 'F') || ('a' <= r && r <= 'f')
}
