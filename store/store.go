package store

import (
	"errors"
)

type Account struct {
	FirstName, LastName string
}

func (a *Account) ChangeName(whichName string, newName string) error {
	switch {
	case whichName == "first":
		a.FirstName = newName
	case whichName == "last":
		a.LastName = newName
	default:
		return errors.New("name does not support more than first and last")
	}

	return nil
}
