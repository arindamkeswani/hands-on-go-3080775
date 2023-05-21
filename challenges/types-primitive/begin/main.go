// challenges/types-primitive/begin/main.go

package main

import (
	"fmt"
)

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global";
func main() {
	// create a local variable "val" and assign it an int value
	var val = 2;
	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v\n", val ,val);
	// print the value of the package-level variable "val"
	printGlobalVal();
	// update the package-level variable "val" and print it again
	updateGlobalVal();
	printGlobalVal();
	// print the pointer address of the local variable "val"
	fmt.Printf("local val address = %v\n", &val)
	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 123;
	fmt.Printf("directly modified pointer = %v\n", val)
}

func printGlobalVal(){
	fmt.Println("global val = ",val )
}

func updateGlobalVal(){
	val = "updated global"
}