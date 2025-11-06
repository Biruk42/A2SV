package services

import (
	"errors"
	"fmt"
	"library_management/models"
	"sync"
	"time"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ReserveBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

type Library struct {
	Books        map[int]models.Book
	Members      map[int]*models.Member
	mutex        sync.Mutex
	ReserveQueue chan ReservationRequest
}

func NewLibrary() *Library {
	lib := &Library{
		Books:        make(map[int]models.Book),
		Members:      make(map[int]*models.Member),
		ReserveQueue: make(chan ReservationRequest, 10),
	}
	go lib.StartReservationWorker() // start the background worker
	return lib
}

func (l *Library) AddBook(book models.Book) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	book.Status = "Available"
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if _, exists := l.Books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(l.Books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book is not borrowed")
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	book.Status = "Available"
	l.Books[bookID] = book
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var available []models.Book
	for _, b := range l.Books {
		if b.Status == "Available" {
			available = append(available, b)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	member, exists := l.Members[memberID]
	if !exists {
		return nil, errors.New("member not found")
	}
	return member.BorrowedBooks, nil
}
//concurrent
type ReservationRequest struct {
	BookID   int
	MemberID int
	Response chan error
}

func (l *Library) ReserveBook(bookID int, memberID int) error {
	resp := make(chan error)
	req := ReservationRequest{
		BookID:   bookID,
		MemberID: memberID,
		Response: resp,
	}
	l.ReserveQueue <- req
	return <-resp
}

func (l *Library) StartReservationWorker() {
	for req := range l.ReserveQueue {
		go l.processReservation(req)
	}
}

func (l *Library) processReservation(req ReservationRequest) {
	l.mutex.Lock()

	book, exists := l.Books[req.BookID]
	if !exists {
		l.mutex.Unlock()
		req.Response <- errors.New("book not found")
		return
	}
	if book.Status != "Available" {
		l.mutex.Unlock()
		req.Response <- errors.New("book not available for reservation")
		return
	}

	book.Status = "Reserved"
	l.Books[req.BookID] = book
	l.mutex.Unlock()

	fmt.Printf("Book %d reserved by member %d\n", req.BookID, req.MemberID)

	go func(bookID, memberID int) {
		time.Sleep(5 * time.Second)
		l.mutex.Lock()
		defer l.mutex.Unlock()
		b := l.Books[bookID]
		if b.Status == "Reserved" {
			b.Status = "Available"
			l.Books[bookID] = b
			fmt.Printf("Reservation expired for book %d\n", bookID)
		}
	}(req.BookID, req.MemberID)

	req.Response <- nil
}
