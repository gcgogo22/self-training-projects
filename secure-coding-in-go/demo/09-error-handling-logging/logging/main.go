package main

import (
	"bytes"
	"fmt"
	"log"
)

/*
By using a buffer, you have more control over when and how logs are displayed or written, which can be useful in cases where you want to capture logs and output them later, rather than immediately printing them. 
*/

func main() {
	var buf bytes.Buffer
	var RoleLevel int

	logger := log.New(&buf, "logger: ", log.Lshortfile)

	fmt.Println("Please enter your user level.")
	fmt.Scanf("%d", &RoleLevel) // <---- example

	switch RoleLevel {
	case 1:
		// Log successful login
		logger.Printf("Login successful.")
		fmt.Print(&buf)
	case 2:
		// Log unsuccessful login
		// When you later call fmt.Print(&buf) or fmt.Print(buf.String()), it prints the content of the buffer to standard output. 

		// bytes.Buffer implements the String() method, so calling fmt.Print(&buf) or fmt.Print(buf.String()) will print the content of the buffer. 
		logger.Printf("Login unsuccessful - Insufficient access level")
		fmt.Print(&buf)
	default:
		// Unspecified error
		logger.Print("Login error.")
		fmt.Print(&buf)
	}
}
