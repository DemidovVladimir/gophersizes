package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"unicode"
)

// Cipher encrypts and decrypts a string.
type Cipher interface {
	Encryption(string) string
	Decryption(string) string
}

// Cipher holds the key used to encrypts and decrypts messages.
type cipher []int

// cipherAlgorithm encodes a letter based on some function.
func (c cipher) cipherAlgorithm(letters string, shift func(int, int) int) string {
	shiftedText := ""
	for _, letter := range letters {
		if !unicode.IsLetter(letter) {
			continue
		}
		shiftDist := c[len(shiftedText)%len(c)]
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		switch {
		case s < 'a':
			s += 'z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(s)
	}
	return shiftedText
}

// Encryption encrypts a message.
func (c cipher) Encryption(plainText string) string {
	return c.cipherAlgorithm(plainText, func(a, b int) int { return a + b })
}

// Decryption decrypts a message.
func (c cipher) Decryption(cipherText string) string {
	return c.cipherAlgorithm(cipherText, func(a, b int) int { return a - b })
}

// NewCaesar creates a new Caesar shift cipher.
func NewCaesar(key int) Cipher {
	return NewShift(key)
}

// NewShift creates a new Shift cipher.
func NewShift(shift int) (c Cipher) {
	if shift < -25 || 25 < shift || shift == 0 {
		return nil
	}
	c = cipher([]int{shift})
	return
}

func main() {

	// This part uses custom crypto algorithm

	// c := NewCaesar(1)
	// fmt.Println("Encrypt Key(01) abcd =>", c.Encryption("ab5cd"))
	// fmt.Println("Decrypt Key(01) bcde =>", c.Decryption("bcde"))
	// fmt.Println()

	// c = NewCaesar(10)
	// fmt.Println("Encrypt Key(10) abcd =>", c.Encryption("abcd"))
	// fmt.Println("Decrypt Key(10) klmn =>", c.Decryption("klmn"))
	// fmt.Println()

	// c = NewCaesar(15)
	// fmt.Println("Encrypt Key(15) abcd =>", c.Encryption("abcd"))
	// fmt.Println("Decrypt Key(15) pqrs =>", c.Decryption("pqrs"))

	// Uses hash typed algos
	fmt.Println("\n----------------Small Message----------------\n")
	message := []byte("Today web engineering has modern apps adhere to what is known as a single-page app (SPA) model.")

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))

	fmt.Println("\n\n----------------Large Message----------------\n")
	message = []byte("Today web engineering has modern apps adhere to what is known as a single-page app (SPA) model. This model gives you an experience in which you never navigate to particular pages or even reload a page.  It loads and unloads the various views of our app into the same page itself. If you've ever run popular web apps like Gmail, Facebook, Instagram, or Twitter, you've used a single-page app. In all those apps, the content gets dynamically displayed without requiring you to refresh or navigate to a different page. React gives you a powerful subjective model to work with and supports you to build user interfaces in a declarative and component-driven way.")

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))
}
