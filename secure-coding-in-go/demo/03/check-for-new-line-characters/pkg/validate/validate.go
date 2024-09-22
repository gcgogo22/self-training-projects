package validate

import "strings"

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (*Validator) ContainsNewLine(input string) bool {
	index := strings.IndexByte(input, '\n') // Find the first occurance of the newline character.
	return index >= 0
}

// http://localhost:8080/containsNewLine?input=Hello%0ATherel