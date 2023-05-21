// packages/basics/main.go
package main

import (
	"fmt"
	"strconv"
	"errors"
	proverbs "github.com/jboursiquot/go-proverbs"
	// "time"
)

func greet() string{
	return "Greet function"
}

func greetWithName(name string) string{
	return "Hello "+name+"!"
}

func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, I am "+name+" and I am "+strconv.Itoa(age)+" years old!"
	return;
}

func divide(a,b int) (int, error){
	if b==0 {
		return 0, errors.New("Cannot divide by 0")
	}

	return a/b, nil
}

func getRandomProverb() string {
	return proverbs.Random().Saying;
}

func main() {
	// fmt.Println(greet());
	// fmt.Println(greetWithName("Arindam"));
	// fmt.Println(greetWithNameAndAge("Arindam", 23));
	// fmt.Println(divide(45, 23));
	// fmt.Println(divide(45, 0));
	fmt.Println(getRandomProverb());
	// fmt.Printf("Today is %s", time.Now().Weekday())
}