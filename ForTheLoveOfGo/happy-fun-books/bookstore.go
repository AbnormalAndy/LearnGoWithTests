package bookstore

import (
	"errors"
	"fmt"
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

func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}

	return result
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d does not exist.", ID)
	}

	return b, nil
}
