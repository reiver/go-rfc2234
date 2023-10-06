package rfc2234

// IsBit returns true if the value of 'r' matches 'BIT' as defined in IETF RFC-2234:
//
//	BIT =  "0" / "1"
func IsBit(r rune) bool {
	return ('0' == r || '1' == r)
}
