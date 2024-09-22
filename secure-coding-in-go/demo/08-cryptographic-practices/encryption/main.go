package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)


// What sets GCM apart is its status as an authenticated cipher mode. Not just a mode for encryption, but also provides authentication for the data, which is a key feature distinguishing it from many other encryption modes. 

// Authentication ensures that the data has not been tampered with and the message is from a trusted source. It guarantees integrity and authenticity by detecting any unauthorized modifications to the encrypted message. 

func encrypt(val []byte, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return aead.Seal(nonce, nonce, val, nil), nil
}

func decrypt(val []byte, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	size := aead.NonceSize()
	if len(val) < size {
		return nil, err
	}

	// Verify the authentication tag and return the decrypted data.
	result, err := aead.Open(nil, val[:size], val[:size], nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func secret() ([]byte, error) {
	// 16 bytes of random secret key
	key := make([]byte, 16)

	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}

func main() {
	secret, err := secret()
	if err != nil {
		log.Fatalf("unable to create secret key: %v", err)
	}

	message := []byte("Welcome to Go Language Secure Coding Practices")
	log.Printf("Message :%s\n", message)

	encrypted, err := encrypt(message, secret)
	if err != nil {
		log.Fatalf("unable to encrypt the data: %v", err)
	}
	log.Printf("Encrypted: %x\n", encrypted)

	decrypted, err := decrypt(encrypted, secret)
	if err != nil {
		log.Fatalf("unable to decrypt the data: %v", err)
	}
	log.Printf("Decrypted: %s\n", decrypted)
}
