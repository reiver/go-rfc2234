package rfc2234

// IsSubDelim returns true if the value of 'r' matches 'sub-delims' as defined in IETF RFC-2234:
//
//	sub-delims  = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
func IsSubDelim(r rune) bool {
	switch r {
	case          '!' , '$' , '&' ,'\'' , '(' , ')' , '*' , '+' , ',' , ';' , '=':
		return true
	default:
		return false
	}
}
