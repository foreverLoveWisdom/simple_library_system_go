package main

func main() {
	myLibrary := Library{}
	myLibrary = myLibrary.AddBook("1984", "George Orwell", 328)
	myLibrary = myLibrary.AddBook("The Hobbit", "J.R.R. Tolkien", 310)
	myLibrary = myLibrary.RemoveBook("The Hobbit")
	myLibrary.DisplayBooks()
}
