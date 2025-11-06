package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	Library *services.Library
}

func NewLibraryController(library *services.Library) *LibraryController {
	return &LibraryController{Library: library}
}

func (c *LibraryController) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			c.addBook(reader)
		case "2":
			c.removeBook(reader)
		case "3":
			c.borrowBook(reader)
		case "4":
			c.returnBook(reader)
		case "5":
			c.listAvailableBooks()
		case "6":
			c.listBorrowedBooks(reader)
		case "7":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *LibraryController) addBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Enter Title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter Author: ")
	author, _ := reader.ReadString('\n')

	book := models.Book{
		ID:     id,
		Title:  strings.TrimSpace(title),
		Author: strings.TrimSpace(author),
	}
	c.Library.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (c *LibraryController) removeBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	err := c.Library.RemoveBook(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Book removed successfully!")
}

func (c *LibraryController) borrowBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to borrow: ")
	bookStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookStr))

	fmt.Print("Enter Member ID: ")
	memberStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberStr))

	if _, exists := c.Library.Members[memberID]; !exists {
		fmt.Print("Enter Member Name: ")
		name, _ := reader.ReadString('\n')
		c.Library.Members[memberID] = &models.Member{ID: memberID, Name: strings.TrimSpace(name)}
	}

	err := c.Library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Book borrowed successfully!")
}

func (c *LibraryController) returnBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to return: ")
	bookStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookStr))

	fmt.Print("Enter Member ID: ")
	memberStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberStr))

	err := c.Library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Book returned successfully!")
}

func (c *LibraryController) listAvailableBooks() {
	books := c.Library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}
	fmt.Println("Available Books:")
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func (c *LibraryController) listBorrowedBooks(reader *bufio.Reader) {
	fmt.Print("Enter Member ID: ")
	memberStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberStr))

	books, err := c.Library.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(books) == 0 {
		fmt.Println("No borrowed books for this member.")
		return
	}

	fmt.Println("Borrowed Books:")
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}
