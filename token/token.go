// Package token defines tokens, their types and functionality for token
// categorization.
//
// TODO: add Token category functions (IsLiteral, ...)
//
// TODO: add ToString for Token
package token

import "github.com/relnod/alp/source"

// Type defines the type of a token.
type Type = uint8

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

// Token defines a token including its type, possibly empty value, and original
// position in the source code.
type Token struct {
	Type  Type
	Value string
	Span  Span
}

// Span defines the span of a token inside a piece of source code.
type Span struct {
	Source source.ID
	Line   source.Line
	ColBeg source.Column
	ColEnd source.Column
}
