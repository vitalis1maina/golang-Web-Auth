package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var key []byte

func main() {
	//use cases
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}
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
func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error in sign message while hashing")
	}
	signature := h.Sum(nil)
	return signature, nil
}
func checksig(msg, sig []byte) (bool, error) {
	newsig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("error in checkSig while getting signature of message: %w", err)
	}
	same := hmac.Equal(newsig, sig)
	return same, nil
}
