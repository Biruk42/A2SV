package main

import (
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	library := services.NewLibrary()

	library.AddBook(models.Book{ID: 1, Title: "Atomic Habit", Author: "Biruk"})
	library.AddBook(models.Book{ID: 2, Title: "Library", Author: "Biruk"})



	controller := controllers.NewLibraryController(library)
	controller.Run()
}
