package password_hash

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Password_test() {

	pwd := getPwd("Enter a password")
	hash := HashAndSalt(pwd)
	fmt.Print("hash_password:", hash, "\n\n\n")

	pwd = getPwd("Enter a password")
	cmp_hash_pwd := getPwd("Enter a hash password")

	pwdMatch := ComparePasswords(string(cmp_hash_pwd), pwd)

	fmt.Print("Passwords Match?", pwdMatch, "\n\n\n")

}

func getPwd(message string) []byte {

	// Prompt the user to enter a password
	fmt.Println(message)

	// We will use this to store the users input
	var pwd string

	// Read the users input
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}

	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(pwd)

}

func HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {

	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}
