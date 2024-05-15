package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:	"Spark Joy",
		Author:	"Marie Kondo",
		Copies:	2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:	"Spark Joy",
		Author:	"Marie Kondo",
		Copies:	2,
	}
	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("Want %d; got %d.", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:	"Spark Joy",
		Author:	"Marie Kondo",
		Copies:	0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("Want error buying from zero copies. Got nil.")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := bookstore.Book{
		ID: 2,
		Title: "The Power of Go: Tools",
	}
	got, err := bookstore.GetBook(catalog, 2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{}
	_, err := bookstore.GetBook(catalog, 999)
	if err == nil {
		t.Fatal("Want error for non-existent ID; got nil.")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
		PriceCents: 4000,
		DiscountPercent: 25,
	}
	want := 3000
	got := bookstore.NetPriceCents(b)
	if want != got {
		t.Errorf("Want %d; got %d.", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
		PriceCents: 4000,
	}
	want := 3000
	err := b.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.PriceCents
	if want != got {
		t.Errorf("Want updated price %d; got %d.", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
		PriceCents: 4000,
	}
	err := b.SetPriceCents(-1)
	if err == nil {
		t.Fatal("Want error setting invalid price -1; got nil.")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	cats := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}
	for _, cat := range cats {
		err := b.SetCategory(cat)
		if err != nil {
			t.Fatal(err)
		}
		got := b.Category()
		if cat != got {
			t.Errorf("Want category %q; got %q.", cat, got)
		}
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := b.SetCategory(999)
	if err == nil {
		t.Fatal("Want error for invalid category; got nil.")
	}
}

