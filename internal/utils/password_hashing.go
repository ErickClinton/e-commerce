package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword recebe uma senha em texto plano e retorna o hash da senha.
func HashPassword(password string) (string, error) {
	// Gerar o hash da senha usando bcrypt.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash recebe uma senha em texto plano e um hash, e verifica se o hash corresponde à senha.
func CheckPasswordHash(password, hash string) bool {
	// Comparar a senha com o hash usando bcrypt.
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
