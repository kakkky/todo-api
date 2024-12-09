package hash

import "golang.org/x/crypto/bcrypt"

func Hash(target string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(target), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func Compare(hashed, target string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(target))
	return err == nil
}
