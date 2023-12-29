// Write a program which can compute the factorial of a given numbers.
// The results should be printed in a comma-separated sequence on a single line.

// Suppose the following input is supplied to the program: 8
// Then, the output should be: 40320

package main

import (
	"fmt"
	"log"
)

func main() {
	var input int
	fmt.Println("Enter a number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal("Please enter a number")
	}
	result, err := fact(input)
	if err != nil {
		log.Fatalf("Error for input %v: %v", input, err)
	}
	fmt.Printf("The factorial of %d is %d", input, result)

}

func fact(n int) (int, error) {
	var f int = 1
	if n == 0 {
		return 1, nil
	}
	if n < 0 {
		return 0, fmt.Errorf("factorial for negativ values is not allowed")
	}

	for i := 1; i <= n; i++ {
		f *= i
	}
	return f, nil

}
