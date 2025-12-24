package crypto

import "golang.org/x/crypto/bcrypt"

type HashCrypto struct{}

func NewHashCrypto() *HashCrypto {
	return &HashCrypto{}
}

func (h *HashCrypto) Hash(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *HashCrypto) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
