package util

import "golang.org/x/crypto/bcrypt"

func CheckPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	return err == nil
}