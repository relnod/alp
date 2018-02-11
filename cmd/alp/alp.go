package main

import (
	"fmt"

	"github.com/relnod/alp/token"
)

func main() {
	t := token.Token{
		Type:  token.Int,
		Value: "123",
	}

	fmt.Println(t)
}
