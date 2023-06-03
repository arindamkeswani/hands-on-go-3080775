// types/maps/initialization/begin/main.go
package main

import "fmt";

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	var authors map[string]author

	// initialize the map with make
	authors = make(map[string]author)

	// add authors to the map
	authors["ak"] = author{name:"Arindam Keswani"}
	authors["ma"] = author{name:"Marcus Aurelius"}

	//alt approach
	// authorsAlt := map[string]author{
	// 	"ak": {name: "Arindam"},
	// 	"ma": {name: "Marcus"},
	// }

	// print the map with fmt.Printf
	fmt.Printf("%#v\n",authors);
	// fmt.Printf("ALT: %#v\n",authorsAlt);

	// read a value from the map with a known key
	fmt.Println(authors["ak"])
	fmt.Println(authors["a"])

	// comma-okay style to prevent empty returns in case on non-existent key
	a, ok := authors["a"]
	fmt.Printf("a = %v, ok = %v\n", a, ok)

	//update value
	
}
