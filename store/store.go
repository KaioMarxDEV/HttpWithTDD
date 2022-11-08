package store

import (
	"errors"
	"fmt"
)

// Account serves as the first abstraction of an user in real world
// it is base for using methods seen below like Stringer()
type Account struct {
	FirstName, LastName string
}

// Employee is a custom type used to implements everything of Account
// and also add more variables and methods like seen below its definition:
type Employee struct {
	Account
	numCred float64
}

func CreateEmployee(first string, last string, credits float64) (*Employee, error) {
	return &Employee{Account{first, last}, credits}, nil
}

// this method is related to objects whom implements account structure
// and changes the variables saved in Account by parameters passed
func (a *Account) ChangeName(newFirst string, newLast string) {
	if newFirst != "" {
		a.FirstName = newFirst
	}
	if newLast != "" {
		a.LastName = newLast
	}

	fmt.Printf("Name changed to %v %v successfully!\n", a.FirstName, a.LastName)
}

// method AddCredits receveis a float and increase the total amount already
// saved into numCred variable inside the implementation of the receiver (Employee)
func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.numCred += amount
		return e.numCred, nil
	}
	return 0.0, errors.New("invalid credit amount")
}

// method RemoveCredits receveis a float and decrease the total amount already
// saved into numCred variable inside the implementation of the receiver (Employee)
func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.numCred {
			return e.numCred, nil
		}
		return 0.0, errors.New("you can't remove more credits than the account has")
	}
	return 0.0, errors.New("you can't remove negative numbers")
}

// method CheckCredits is used to verify the total amount in real time inside
// the variable numCred from an object that implements Employee
func (e Employee) CheckCredits() float64 {
	return e.numCred
}

// method Stringer formats a string with variable values found in Account implementation
// to use this method the object needs to be an instance of Account struct
// Employee is an example of Account struct from embedding structs.
func (e Employee) Stringer() string {
	return fmt.Sprintf("Account data Name: %v %v Credits: %.2f", e.FirstName, e.LastName, e.numCred)
}
