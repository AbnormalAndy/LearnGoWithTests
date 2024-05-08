package bookstore

import (
	"errors"
)

type Book struct {
	ID	int
	Title	string
	Author	string
	Copies	int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("No copies left.")
	}
	b.Copies--
	return b, nil 
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog map[int]Book, ID int) Book {
	return catalog[ID]
}
