// Package token defines tokens, their types and functionality for token
// categorization.
//
// TODO: add Token category functions (IsLiteral, ...)
//
// TODO: add ToString for Token
package token

import (
	"fmt"

	"github.com/relnod/alp/source"
)

// Type defines the type of a token.
type Type uint8

// Token types
const (
	Invalid Type = iota

	EOF
	Comment

	literalBegin
	Ident // Identifier
	Int   // integer literal
	Float // floating-point literal
	literalEnd

	operatorBegin
	Assign    // =
	Semicolon // ;
	Colon     // :

	Add // +
	Sub // -
	Mul // *
	Div // /
	Mod // %
	operatorEnd

	keywordBegin
	Let // keyword: let
	keywordEnd
)

// types defines the string representation of token types
var types = [...]string{
	Invalid: "Invalid",

	EOF:     "EOF",
	Comment: "Comment",

	Ident: "Ident",
	Int:   "Int",
	Float: "Float",

	Assign:    "=",
	Semicolon: ";",
	Colon:     ":",

	Add: "+",
	Sub: "-",
	Div: "/",
	Mod: "%",

	Let: "let",
}

// String converts a token type to it's string representation.
func (t Type) String() string {
	s := ""
	if t < Type(len(types)) {
		s = types[t]
	}

	if s == "" {
		s = fmt.Sprintf("Unkown Token: %d", t)
	}

	return s
}

// Token defines a token including its type, possibly empty value, and original
// position in the source code.
type Token struct {
	Type  Type
	Value string
	Span  Span
}

// String converts a token to a string.
func (t Token) String() string {
	return fmt.Sprintf("{ Type: %s, Value: \"%s\", Span: %s }",
		t.Type, t.Value, t.Span)
}

// Span defines the span of a token inside a piece of source code.
type Span struct {
	Source source.Source
	Line   source.Line
	ColBeg source.Column
	ColEnd source.Column
}

// String converts a span to its string representation
//
// TODO: convert Source to a more meaningfull string
func (s Span) String() string {
	sourceName := ""
	if s.Source != nil {
		sourceName = s.Source.Name()
	}

	return fmt.Sprintf("{ Source: %s, Line: %d, ColBeg: %d, ColEnd: %d }",
		sourceName, s.Line, s.ColBeg, s.ColEnd)
}
