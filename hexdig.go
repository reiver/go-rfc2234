package rfc2234

// IsHexDigit returns true if value of 'r' matches 'ALPHA' as defined in IETF RFC-2234:
//
//	HEXDIG =  DIGIT / "A" / "B" / "C" / "D" / "E" / "F"
//	
//	DIGIT          =  %x30-39
//	                       ; 0-9
func IsHexDig(r rune) bool {
	return ('0' <= r && r <= '9') || ('A' <= r && r <= 'F')
}
