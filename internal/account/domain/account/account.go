package account

import "time"

// Account is an entity for the account information registered by the User.
type Account struct {
	ID           uint64
	email        string
	password     string
	registerDate *time.Time
	Status       Status
}

// IsActive is to check whether the status of the account is normal.
func (acc *Account) IsActive() bool {
	return acc.Status == ACTIVE
}

// IsValid is to check email and password are valid
func (acc *Account) IsValid() bool {
	email := NewEmailSpecification()
	password := NewPasswordSpecification()

	accountSpec := email.And(password)

	return accountSpec.IsSatisfiedBy(*acc)
}

// GetEmail is a getter
func (acc *Account) GetEmail() string {
	return acc.email
}

// GetPassword is a getter
func (acc *Account) GetPassword() string {
	return acc.password
}
