package main

import (
	"fmt"
	"simple_library_system_go/library"
)

func main() {
	ebook := &library.Ebook{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"}
	fmt.Println(ebook.GetTitle(), ebook.GetAuthor(), ebook.GetType())
}
