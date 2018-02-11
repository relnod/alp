package main

import (
	"fmt"

	"github.com/relnod/alp/source"
	"github.com/relnod/alp/token"
)

func main() {
	fs := source.NewFileSource("test/bla.alp")
	t1 := token.Token{
		Type:  token.Int,
		Value: "123",
		Span: token.Span{
			Source: &fs,
			Line:   10,
			ColBeg: 5,
			ColEnd: 10,
		},
	}
	r, err := fs.GetReader()
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 10)
	r.Read(buffer)
	fmt.Println(string(buffer))

	fmt.Println(t1)

	es := source.NewEvalSource("123", "let a: int = 2")
	t2 := token.Token{
		Type:  token.Int,
		Value: "123",
		Span: token.Span{
			Source: &es,
			Line:   10,
			ColBeg: 5,
			ColEnd: 10,
		},
	}

	fmt.Println(t2)

	r2, _ := fs.GetReader()

	r2.Read(buffer)
	fmt.Println(string(buffer))
}
