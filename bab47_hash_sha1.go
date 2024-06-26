package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func hashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	// * creating different salted text
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	encrypt := sha.Sum(nil)
	return fmt.Sprintf("%x", encrypt), salt
}

func main() {
	text := "this is secret"
	fmt.Printf("Original : %s\n\n", text)

	var hashed1, salt1 = hashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)

	var hashed2, salt2 = hashUsingSalt(text)
	fmt.Printf("hashed 2 : %s\n\n", hashed2)

	var hashed3, salt3 = hashUsingSalt(text)
	fmt.Printf("hashed 3 : %s\n\n", hashed3)

	_, _, _ = salt1, salt2, salt3
}
