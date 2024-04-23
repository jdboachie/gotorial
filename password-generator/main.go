package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func generatePassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[]:;,.?"
	charsetLength := big.NewInt(int64(len(charset)))

	password := make([]byte, length)
	for i := range password {
		n, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}

	return string(password), nil
}

func main() {
	length := flag.Int("length", 12, "Length of the password")
	flag.Parse()

	password, err := generatePassword(*length)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(password)
}