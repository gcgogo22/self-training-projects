package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/blake2s"
)

func main() {
	md5Hash := md5.New()
	sha256Hash := sha256.New()
	h_blake2s, _ := blake2s.New256(nil)

	io.WriteString(md5Hash, "Welcome to Go Language Secure Coding Practices")
	io.WriteString(sha256Hash, "Welcome to Go Language Secure Coding Practices")
	io.WriteString(h_blake2s, "Welcome to Go Language Secure Coding Practices")

	// Sum with passed nil returns the final MD5 hash value as a byte slice.
	// During hashing, you progressively add data to the hash methods like Write(). After all data is written, the Sum() method is used to get the final result (the hash).  

	// The hashing functino returns a slice of bytes. 
	// The %x format specifier is used to convert each byte of the slice into its hexdecimal representation. 

	// The output of the hash function is a binary value (a fixed-size sequence bits), which is typically represented as a slice of bytes in programming language like Go. 

	// The result of a hash function is fundamentally binary data, which is why it is represented as a sequence of bytes, not single number. The slice of bytes allows you to handle the binary output efficiently.


	fmt.Printf("MD5    		:%x\n", md5Hash.Sum(nil))
	fmt.Printf("SHA256 		:%x\n", sha256Hash.Sum(nil))
	fmt.Printf("Blake2s-256	:%x\n", h_blake2s.Sum(nil))
}
