package password

import (
	"goblong/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// Generate from password ret urns the  bcry has of the password
func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogError(err)

	return string(bytes)
}

// Compare as bcrypt hashed password with its possible plaintext equivalent.
// Returns true success
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	logger.LogError(err)
	return err == nil
}

// Check string is hashed
func IsHashed(str string) bool {
	return len(str) == 60
}
