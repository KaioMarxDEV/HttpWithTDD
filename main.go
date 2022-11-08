package main

import (
	"fmt"
	"test/store"
)

func main() {
	newUser := store.Account{FirstName: "john", LastName: "doe"}
	newEmployee := store.Employee{}
	var newFirst, newLast string
	var amountToAdd, amountToRemove float64

	fmt.Println("\n==== WELCOME to STORE. accounts session ====")
	fmt.Println("The program auto determined a name for your account")

	fmt.Println("To custom it, please digits your first name:")
	fmt.Scan(&newFirst)
	fmt.Println("Now the your last name:")
	fmt.Scan(&newLast)

	newUser.ChangeName(newFirst, newLast)

	fmt.Println("==== WELCOME to STORE. Employee session ====")
	fmt.Println("Here you can Increase, Decrease and Check your credits")

	fmt.Println("Please digit how much you want to add:")
	fmt.Scan(&amountToAdd)
	newEmployee.AddCredits(amountToAdd)

	fmt.Printf("\nYour Amount now is %.2f\n", newEmployee.CheckCredits())

	fmt.Println("Please digit how much you want to REMOVE:")
	fmt.Scan(&amountToRemove)
	newEmployee.RemoveCredits(amountToRemove)

	fmt.Printf("\nYour Amount now is %.2f\n", newEmployee.CheckCredits())
	fmt.Println("Progam is now finished...")
}
