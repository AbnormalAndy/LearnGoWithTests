package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID		int
	Title		string
	Author		string
	Copies		int
	PriceCents	int
	DiscountPercent	int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("No copies left.")
	}
	b.Copies--
	return b, nil 
}

type Catalog map[int]Book

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
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

func NetPriceCents(b Book) int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("Negative price %d.", price)
	}
	b.PriceCents = price
	return nil
}
