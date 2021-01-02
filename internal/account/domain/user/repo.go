package user

// IUserRepository is an UserRepository interface
type IUserRepository interface {
	GetUserByAccountID() (*User, error)

	Save(user User) (*User, error)
}
