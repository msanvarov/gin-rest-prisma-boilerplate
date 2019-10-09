package utils

import "golang.org/x/crypto/bcrypt"

// hash password and return either error or string in bytes
// note: the higher the hashing cost, the more computing power is required resulting in slower server response times
func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// check the password against the hashed password to validate for similarities
func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}