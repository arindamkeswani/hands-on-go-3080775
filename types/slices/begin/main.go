// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string the make function
	names := make([]string, 0)

	fmt.Println(names);
	// append 3 names to the slice
	names = append(names, "John")
	names = append(names, "Jane")
	names = append(names, "Mary")
	
	// print the slice
	fmt.Println(names);

	// slice the slice using syntax slice[low:high]
	fmt.Println(names[1:3])// [Jane Mary] (high-bound is always non-incl)
	fmt.Println(names[1:])// [Jane Mary] (no high-bound)
	fmt.Println(names[:1])// [John] (no low bound)
	fmt.Println(names[:])// [John Jane Mary]
}
