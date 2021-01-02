package account

// IAccountRepository is an AccountRepository interface
type IAccountRepository interface {
	GetAccountByID() (*Account, error)
	GetAccountByEmail() (*Account, error)

	Save(account Account) (*Account, error)
}
