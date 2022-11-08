package store

import "fmt"

// Account serves as the first abstraction of an user in real world
// it is base for using methods seen below like Stringer()
type Account struct {
	FirstName, LastName string
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

// Employee is a custom type used to implements everything of Account
// and also add more variables and methods like seen below its definition:
type Employee struct {
	Account
	numCred float64
}

// method AddCredits receveis a float and increase the total amount already
// saved into numCred variable inside the implementation of the receiver (Employee)
func (e *Employee) AddCredits(amount float64) float64 {
	e.numCred += amount

	return e.numCred
}

// method RemoveCredits receveis a float and decrease the total amount already
// saved into numCred variable inside the implementation of the receiver (Employee)
func (e *Employee) RemoveCredits(amount float64) float64 {
	e.numCred -= amount

	return e.numCred
}

// method CheckCredits is used to verify the total amount in real time inside
// the variable numCred from an object that implements Employee
func (e Employee) CheckCredits() float64 {
	return e.numCred
}

// method Stringer formats a string with variable values found in Account implementation
// to use this method the object needs to be an instance of Account struct
// Employee is an example of Account struct from embedding structs.
func (e Account) Stringer() string {
	return fmt.Sprintf("Your Account data, first Name: %v\tlast Name: %v", e.FirstName, e.LastName)
}
