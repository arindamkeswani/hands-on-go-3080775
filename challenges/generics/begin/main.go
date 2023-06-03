// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

func print[T any](v T) {
	fmt.Println(v)
}

// Part 2 sum function refactoring
type numeric interface{
	constraints.Integer | constraints.Float
}

// sum sums a slice of any type
func sum[T numeric](numbers ...T) T {
	var result T
	for _, n := range numbers {
		result += n;
	}
	return result
}

func main() {
	// call non-generic print functions
	// printString("Hello")
	// printInt(42)
	// printBool(true)

	// call generic printAny function for each value above
	print("Hello")
	print(42)
	print(true)
	// call sum function
	// fmt.Println("result", sum([]interface{}{1, 2, 3}))
	fmt.Println("result", sum(1.2,2.1,3.1,4.1))

	// call generics sumAny function
}
