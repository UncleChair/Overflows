package bcrypt

import (
	"overflows/internal/service"

	"golang.org/x/crypto/bcrypt"
)

type sBcrypt struct{}

func init() {
	service.RegisterBcrypt(New())
}

func New() *sBcrypt {
	return &sBcrypt{}
}
func (b *sBcrypt) Generate(password string) (hashedPassword string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword = string(hashed)
	return
}
func (b *sBcrypt) Compare(hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
