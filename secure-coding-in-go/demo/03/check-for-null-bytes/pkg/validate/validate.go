package validate

import "strings"

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (iv *Validator) ContainsNullByte(input string) bool {
	if input == "" {
		return false
	}
	return strings.Contains(input, "\x00") // Check if string contains null byte.
}

/*
Input a null value in the http request
http://localhost:8080/containsnullbyte?input=%00

Need to encode null byte in the utf-8 format in order to use it in the http request. 
*/
