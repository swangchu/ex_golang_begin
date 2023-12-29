// With a given integral number n, write a program to generate a map
// that contains (i, i*i) such that is an integral number between 1 and
// n (both included), and then the program should print the map with
// representation of the value

// Suppose the following input is supplied to the program: 8

// Then, the output should be: map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	fmt.Println("Enter a number: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatal("There is an error")
	}
	result := square(n)
	fmt.Printf(" %v", result)
}

func square(n int) map[int]int {
	var sq = make(map[int]int, n)
	for i := 1; i <= n; i++ {
		sq[i] = i * i
	}
	return sq
}
