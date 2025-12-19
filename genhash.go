package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password123!"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println("PASSWORD:", password)
	fmt.Println(string(hash))  #when updating the pw,escape every $ with \
}
