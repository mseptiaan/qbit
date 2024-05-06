package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type Hashing struct{}

func NewHashingPassword() *Hashing {
	return &Hashing{}
}

func (h Hashing) HashPassword(password string) string {
	pw := []byte(password)
	hashedPw, _ := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	return string(hashedPw)
}

func (h Hashing) ComparePassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
