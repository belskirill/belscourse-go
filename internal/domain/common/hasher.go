package common

type PasswordHasher interface {
	Hash(password string, cost int) (string, error)
	Compare(hash, password string) error
}
