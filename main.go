package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "passworD"
	hash, _ := hashPassword(password)

	fmt.Println("Password is:", password)
	fmt.Println("Hash       :", hash)

	match := checkPasswordHash(password, hash)
	fmt.Println("Is match   :", match)
}
