package account

import (
	"regexp"
	"unicode"
)

type Specification interface {
	IsSatisfiedBy(Account) bool
	And(Specification) Specification
	Or(Specification) Specification
	Not() Specification
	Relate(Specification)
}

// Invoice BaseSpecification
type BaseSpecification struct {
	Specification
}

// Check specification
func (self *BaseSpecification) IsSatisfiedBy(acc Account) bool {
	return false
}

// Condition AND
func (self *BaseSpecification) And(spec Specification) Specification {
	a := &AndSpecification{
		self.Specification, spec,
	}
	a.Relate(a)
	return a
}

// Condition OR
func (self *BaseSpecification) Or(spec Specification) Specification {
	a := &OrSpecification{
		self.Specification, spec,
	}
	a.Relate(a)
	return a
}

// Condition NOT
func (self *BaseSpecification) Not() Specification {
	a := &NotSpecification{
		self.Specification,
	}
	a.Relate(a)
	return a
}

// Relate to specification
func (self *BaseSpecification) Relate(spec Specification) {
	self.Specification = spec
}

/////

// AndSpecification
type AndSpecification struct {
	Specification
	compare Specification
}

// Check specification
func (self *AndSpecification) IsSatisfiedBy(acc Account) bool {
	return self.Specification.IsSatisfiedBy(acc) && self.compare.IsSatisfiedBy(acc)
}

/////

// OrSpecification
type OrSpecification struct {
	Specification
	compare Specification
}

// Check specification
func (self *OrSpecification) IsSatisfiedBy(acc Account) bool {
	return self.Specification.IsSatisfiedBy(acc) || self.compare.IsSatisfiedBy(acc)
}

// NotSpecification
type NotSpecification struct {
	Specification
}

// Check specification
func (self *NotSpecification) IsSatisfiedBy(acc Account) bool {
	return self.Specification.IsSatisfiedBy(acc)
}

type EmailSpecification struct {
	Specification
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (spec *EmailSpecification) IsSatisfiedBy(acc Account) bool {
	email := acc.GetEmail()
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}

func NewEmailSpecification() Specification {
	es := &EmailSpecification{&BaseSpecification{}}
	es.Relate(es)
	return es
}

type PasswordSpecification struct {
	Specification
}

func (spec *PasswordSpecification) IsSatisfiedBy(acc Account) bool {
	minLength := 8
	isNumber := false
	isUpper := false

	pw := acc.GetPassword()

	for _, c := range pw {
		switch {
		case unicode.IsNumber(c):
			isNumber = true
		case unicode.IsUpper(c):
			isUpper = true
		}
	}

	return len(pw) >= minLength && isNumber && isUpper
}

func NewPasswordSpecification() Specification {
	ps := &EmailSpecification{&BaseSpecification{}}
	ps.Relate(ps)
	return ps
}
