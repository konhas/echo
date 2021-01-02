package user

import (
	"github.com/biter777/countries"
	"github.com/ttacon/libphonenumber"
)

// Phone is a Value Object that represents phone information in User.
type Phone struct {
	CountryCode     int16
	FormattedNumber string
}

// New is a function that initializes Phone objects.
func New(data string) (Phone, error) {
	pn := Phone{}
	country := countries.Korea

	num, err := libphonenumber.Parse(data, country.Alpha2())
	if err != nil {
		return pn, err
	}

	pn.FormattedNumber = libphonenumber.Format(num, libphonenumber.NATIONAL)

	return pn, nil
}

// GetPhoneNumber is a Getter function that return formatted-number in a Phone Objects
func (p *Phone) GetPhoneNumber() string {
	return p.FormattedNumber
}
