package utils

import "golang.org/x/crypto/bcrypt"

// IPasswords interface.
type IPasswords interface {
	EncryptPassword(password string) (string, error)
	CheckPassword(password string, hash string) bool
}

// EncryptPassword method hashes the inputted password and returns whether or not the operation was successfull.
// Remark: the higher the hashing cost, the more computing power is required resulting in slower server response times
func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPassword method checks the inputted password against the hashed password to validate their hashed equivalence.
func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
