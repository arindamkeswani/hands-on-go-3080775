// challenges/types-composite/begin/main.go
package main

import (
	// "fmt"
	"fmt"

	// "github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title  string
	author author
}

// define a library type to track a list of books
// key: string, value: slice (list) of books
type library map[string][]book

// define addBook to add a book to the library
func (lib library) addBook(b book) {
	lib[b.author.name] = append(lib[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (lib library) lookupByAuthor(authorName string) []book {
	return lib[authorName]
}

func main() {
	// create a new library
	lib := library{}

	// add 2 books to the library by the same author
	ma := author{name: "Marcus"}

	// add 1 book to the library by a different author
	lib.addBook(book{
		title:  "Meditations",
		author: ma,
	})

	lib.addBook(book{
		title:  "Meditations 2",
		author: ma,
	})

	lib.addBook(book{
		title: "Autobio",
		author: author{
			name: "Arindam",
		},
	})

	// dump the library with spew
	// spew.Dump(lib)

	// lookup books by known author in the library
	testLookup := lib.lookupByAuthor("Arindamm");

	// print out the first book's title and its author's name
	if len(testLookup) != 0 {
		b := testLookup[0]
		fmt.Printf("Book title: %v, Author: %v\n",b.title, b.author.name);
	} else{
		fmt.Printf("No books found by searched author\n");
	}

}
