package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	//use cases
	password := "vi&%$#@323vead^!9089u"
	hashedPass, err := hashedPassword(password)
	err = comparePassword(password, hashedPass)
	if err != nil {
		log.Fatalln("Your are not loggedin")
	}
	log.Println("logged in!")
}

func hashedPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating hashed password from bycrpt:%w", err)
	}
	return bs, nil
}
func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid/Wrong Password:%w", err)
	}
	return nil
}
