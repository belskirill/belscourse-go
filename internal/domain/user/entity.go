package user

type Role string

const (
	UserRole  Role = "user"
	AdminRole Role = "admin"
	OwnerRole Role = "owner"
)

type UserCreate struct {
	Username     string
	Email        string
	PasswordHash string
	Password     string
}

type UserBase struct {
	ID       int64
	Username string
	Email    string
}

type UserWithPassword struct {
	ID           int64
	Username     string
	Email        string
	Password     string
	HashPassword string
}
