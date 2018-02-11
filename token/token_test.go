package token_test

import (
	"testing"

	"github.com/relnod/alp/token"
)

func TestTokenTypeToString(t *testing.T) {
	t1 := token.Int
	if t1.String() != "Int" {
		t.Error()
	}

	t2 := token.Type(3)
	if t2.String() != "Unkown Token: 3" {
		t.Error()
	}

	t3 := token.Type(255)
	if t3.String() != "Unkown Token: 255" {
		t.Error()
	}
}
