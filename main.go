package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main()  {
	plainPassword := "pass123"
	fmt.Println("Plain password: ", plainPassword)

	// hash password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(plainPassword),10)
	sHash := string(bytes)
	fmt.Println("Hash password: ", sHash)

	entryPassword := "pass123445"
	fmt.Println("Entry password: ", entryPassword)
	err := bcrypt.CompareHashAndPassword([]byte(sHash), []byte(entryPassword))

	fmt.Println("Password sama ?  ", err == nil)
}