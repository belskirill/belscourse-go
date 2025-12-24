package user

type PasswordService interface {
	HashService(password string, cost int) (string, error)
	CompareService(hash, password string) error
}
