package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	nums := []int{7, 2, 8, -9, 4}

	c := make(chan int)
	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)

	x, y := <-c, <-c
	fmt.Printf("\nx:%v\ty:%v\tTOTAL: %d", x, y, x+y)
}
