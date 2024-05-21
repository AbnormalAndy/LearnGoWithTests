package mytypes

import (
	"strings"
)

type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyBuilder strings.Builder

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}
