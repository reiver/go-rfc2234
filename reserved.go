package rfc2234

// IsReserved returns true if value of 'r' matches 'reserved' as defined in IETF RFC-2234:
//
//	reserved    = gen-delims / sub-delims
//	
//	gen-delims  = ":" / "/" / "?" / "#" / "[" / "]" / "@"
//	
//	sub-delims  = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
func IsReserved(r rune) bool {
	return IsGenDelim(r) || IsSubDelim(r)
}
