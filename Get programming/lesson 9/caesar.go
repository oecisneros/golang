package main

import (
	"fmt"
)

type cipher func(rune) rune

func isLetter(r rune) bool {
	return (r >= 'A' && r <= 'z')
}

func isUpperCase(r rune) bool {
	return (r >= 'A' && r <= 'Z')
}

func isLowerCase(r rune) bool {
	return (r >= 'a' && r <= 'z')
}

// 62 A, 94 a
func toUpper(r rune) rune {
	if isLowerCase(r) {
		return r - 'a' + 'A'
	}
	return r
}

func rot13Cipher(r rune) rune {
	if isLetter(r) {
		c := r + 13
		if toUpper(r)+13 > 'Z' {
			c -= 26
		}
		return c
	}
	return r
}

func caesarCipher(r rune) rune {
	if isLetter(r) {
		c := r - 3
		if toUpper(r)-3 < 'A' {
			c += 26
		}
		return c
	}
	return r
}

func encrypt(t string, c cipher) string {
	buffer := []rune{}
	for _, r := range t {
		buffer = append(buffer, c(r))
	}
	return string(buffer)
}

// go run caesar.go
func main() {
	cipherText := "L fdph, L vdz, L frqtxhuhg"
	plainText := encrypt(cipherText, caesarCipher)

	fmt.Println("Decipher the folowing quote from Julius Caesar:")
	fmt.Println("Ciphertext:", cipherText)
	fmt.Println("Plaintext:", plainText)

	plainText = "Hola Estación Espacial Internacional"
	cipherText = encrypt(plainText, rot13Cipher)

	fmt.Println("Plaintext:", plainText)
	fmt.Println("Ciphertext:", cipherText)
	fmt.Println("Plaintext:", encrypt(cipherText, rot13Cipher))
}
