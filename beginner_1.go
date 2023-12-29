/*
Write a program which will find all such numbers which are divisible by 7
but are not a multiple of 5, between 2000 and 3200 (both included). The
numbers obtained should be printed in a comma-separated sequence on a single line.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numbers []string

	for i := 2000; i <= 3200; i++ {

		if i%7 == 0 && i%5 != 0 {
			numbers = append(numbers, strconv.Itoa(i))
		}
	}
	n := strings.Join(numbers, ",")
	fmt.Println(n)
}

