// flow-control/loop-range/begin/main.go
package main

import "fmt"

func main() {
	// initialize a slice of ints
	nums := []int{100,300,200}

	// use for-range to iterate over the slice and print each value
	//i=index, n=current value
	for i, n := range nums{
		fmt.Println(i, n);
	}

	// declare a map of strings to ints
	m := map[string]int{"one": 1, "two": 2}

	// use for-range to iterate over the map and print each key/value pair
	for key, value:= range m{
		fmt.Println(key, value);
	}
}
