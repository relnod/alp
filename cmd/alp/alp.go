package main

import (
	"fmt"

	"github.com/relnod/alp/source"
	"github.com/relnod/alp/token"
)

func main() {
	s := source.NewFileSource("bla.alp")

	t := token.Token{
		Type:  token.Int,
		Value: "123",
		Span: token.Span{
			Source: &s,
			Line:   10,
			ColBeg: 5,
			ColEnd: 10,
		},
	}

	fmt.Println(t)
}
