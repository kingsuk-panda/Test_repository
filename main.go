package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "mySecretPassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating hash:", err)
		return
	}
	fmt.Println("Hashed password:", string(hashedPassword))
	newPassword := "mySecretPassword"
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(newPassword))
	if err != nil {
		fmt.Println("Password does not match hashed password.")
		return
	}
	fmt.Println("Password matches hashed password.")
}
