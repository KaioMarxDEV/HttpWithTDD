package main

import (
	"fmt"
	"test/store"
)

func main() {
	bruce, _ := store.CreateEmployee("Bruce", "Lee", 500)
	fmt.Println(bruce.CheckCredits())

	credits, err := bruce.AddCredits(250)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	_, err = bruce.RemoveCredits(2500)
	if err != nil {
		fmt.Println("Can't withdraw or overdrawn!", err)
	}

	bruce.ChangeName("Mark", "Zuck")

	fmt.Println(bruce)
}
