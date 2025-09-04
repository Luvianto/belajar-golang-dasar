package commonutils

import "golang.org/x/crypto/bcrypt"

func Encrypt(text *string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
