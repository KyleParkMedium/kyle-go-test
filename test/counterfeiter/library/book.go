package library

import (
	"fmt"
	"strings"
)

type Book struct {
	GUID string

	ISBN string
	Details
}

type Details struct {
	Author string
	Name   string
}

type BookService struct {
	Book    BookStore
	Details BookDetailsStore
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate .

//counterfeiter:generate . BookStore
type BookStore interface {
	Create(isbn string, details Details) (Book, error)
}

//counterfeiter:generate . BookDetailsStore
type BookDetailsStore interface {
	Find(isbn string) (Details, error)
}

func (b *BookService) Create(isbn string) (Book, error) {
	details, err := b.Details.Find(isbn)
	if err != nil {
		return Book{}, fmt.Errorf("details store %v: %w", isbn, err)
	}

	book, err := b.Book.Create(strings.ToUpper(isbn), details)
	if err != nil {
		return Book{}, fmt.Errorf("book store %v: %w", isbn, err)
	}

	return book, nil
}
