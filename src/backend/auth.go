package main

import (
	"golang.org/x/crypto/bcrypt"
)

// Método para transformar a senha em uma senha criptografada - MAIOR SEGURANÇA
func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// Método para comparar a string com a hash armazenada
func checkPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}