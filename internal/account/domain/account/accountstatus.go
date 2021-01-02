package account

// Status is a Value Object for the user account status.
type Status int8

// Status type
const (
	ACTIVE = iota
	WITHDRAWN
	BANNED
)

func (s Status) String() string {
	return [...]string{"active", "withdrawn", "banned"}[s]
}
