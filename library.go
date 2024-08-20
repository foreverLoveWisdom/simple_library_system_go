package main

// Let's apply the questions I mentioned earlier to the current library file.

// **Multiple implementations**:

// * Do you need to support different types of books (e.g., e-books, audiobooks, physical books)?
// * Are there multiple ways to display books (e.g., grid view, list view, table view)?
// * Are there different sources for book data (e.g., database, API, file system)?

// If you answer "yes" to any of these questions, an interface might be necessary to define the contract for different implementations.

// **Decoupling**:

// * Is the `Library` struct tightly coupled to the `Book` struct?
// * Do you want to make it easier to switch from one book storage system to another (e.g., from a slice to a database)?
// * Are there any components that rely heavily on the `Library` struct, making it difficult to test or maintain?

// If you answer "yes" to any of these questions, an interface might help decouple the dependent components.

// **Abstraction**:

// * Are there complex business rules or algorithms related to book management that need to be hidden from the rest of the system?
// * Do you need to provide a simplified API for external developers or users to interact with the library?
// * Are there any components that require a high level of customization or flexibility in terms of book management?

// If you answer "yes" to any of these questions, an interface might help abstract away the complex implementation details.

// **Testability**:

// * Are there any components that are difficult to test due to tight coupling or complex dependencies in the `Library` struct?
// * Do you want to make it easier to write unit tests or integration tests for the library system?
// * Are there any components that require mock implementations or stubs to test effectively?

// If you answer "yes" to any of these questions, an interface might help improve the testability of the system.

// Based on the current implementation, it seems that the `Library` struct is relatively simple and doesn't have many dependencies. However, if you plan to add more features or complexity to the library system, introducing an interface might be beneficial to ensure maintainability, scalability, and testability.

// Do you have any specific plans or requirements for the library system that might necessitate the introduction of an interface? I'd be happy to help you explore those further!

import "fmt"

// Book struct represents a book with a title, author, and pages.
type Book struct {
	Title  string
	Author string
	Pages  int
}

// Library struct represents a collection of books.
type Library struct {
	Books []Book
}

// AddBook adds a new book to the library.
func (l Library) AddBook(title, author string, pages int) Library {
	book := Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
	l.Books = append(l.Books, book)
	return l
}

// RemoveBook removes a book from the library by its title.
func (l Library) RemoveBook(title string) Library {
	for i, book := range l.Books {
		if book.Title == title {
			l.Books = l.Books[:i+copy(l.Books[i:], l.Books[i+1:])]
			break
		}
	}
	return l
}

// DisplayBooks prints out all the books in the library.
func (l Library) DisplayBooks() {
	for _, book := range l.Books {
		fmt.Printf("Title: %s, Author: %s, Pages: %d\n", book.Title, book.Author, book.Pages)
	}
}
