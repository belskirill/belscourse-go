package user

type Role string

const (
	UserRole  Role = "user"
	AdminRole Role = "admin"
	OwnerRole Role = "owner"
)

type User struct {
	ID           int64
	Username     string
	Email        string
	PasswordHash string
}
