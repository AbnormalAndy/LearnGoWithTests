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

type StringUppercaser struct {
	Contents strings.Builder
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}
