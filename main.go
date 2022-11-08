package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

func (merr *MyError) Error() string {
	return fmt.Sprintf("When: %v, What: %v", merr.when, merr.what)
}

func Run() error {
	return &MyError{when: time.Now(), what: "It didn't work out"}
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
