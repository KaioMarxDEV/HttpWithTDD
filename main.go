package main

import (
	"fmt"
	"test/store"
)

func main() {
	a := store.Account{}
	var choice, newName string

	fmt.Println("This is Online Store Program...to manage accounts")

	fmt.Println("Input your first name:")
	fmt.Scan(&a.FirstName)
	fmt.Println("Input your last name:")
	fmt.Scan(&a.LastName)

	fmt.Println("Input where you want to change: [first or last]")
	fmt.Scan(&choice)

	// verifying is updating first or last name...
	choicePossible := [2]string{"first", "last"}
	for _, c := range choicePossible {
		if c == choice {
			break
		}

		fmt.Println("You need to choose just one of the possibles: [first or last]")
		fmt.Scan(&choice)
	}

	fmt.Println("What will be your new name: [digits]")
	fmt.Scan(&newName)

	err := a.ChangeName(choice, newName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\naccount data: %v ---- %v", a.FirstName, a.LastName)
}
