package rfc2234

// IsGenDelim returns true if the value of 'r' matches 'sub-delims' as defined in IETF RFC-2234:
//
//	gen-delims  = ":" / "/" / "?" / "#" / "[" / "]" / "@"
func IsGenDelim(r rune) bool {
	switch r {
	case          ':' , '/' , '?' , '#' , '[' , ']' , '@':
		return true
	default:
		return false
	}
}
