package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			// return x position inside of s when s[i] or v are equal to x
			return i
		}
	}
	return -1
}

func main() {
	sliceInts := []int{10, 20, 14, -2}
	fmt.Println(Index(sliceInts, 20))

	sliceStrings := []string{"hello", "world", "test", "kaio"}
	fmt.Println(Index(sliceStrings, "test"))
}
