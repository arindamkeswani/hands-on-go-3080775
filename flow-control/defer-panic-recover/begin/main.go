// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls (treated like finally from try, catch, finally)
	//executed in LIFO order
	defer cleanup("first")
	defer cleanup("second")

	fmt.Println("Working in main");

	// defer recovery
	defer func(){
		if err:= recover();err!=nil {
			fmt.Println("recovered from panic:", err)
		}
	}()

	// panic
	panic("Server error.")
}
