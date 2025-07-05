package hash

import "golang.org/x/crypto/bcrypt"

type Hash struct{}

func (h *Hash) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.MaxCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func (h *Hash) CompareHashAndPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
