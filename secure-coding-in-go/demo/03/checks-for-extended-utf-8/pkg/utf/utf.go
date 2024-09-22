package utf

import "unicode/utf8"

type Utf struct{}

func NewUtf() *Utf {
	return &Utf{}
}

func (*Utf) InputIsValidUtf8(input string) bool {
	return utf8.ValidString(input)
}

/*
URL encoding converts characters that may not be safely transmitted in URLs into a % followed by two hexadecimal digits. 

For this one, input a string containing non-encoded character returns false.
*/
