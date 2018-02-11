package source

import (
	"io"
)

// Type represents the type of a source.
type Type uint8

// Source types
const (
	File Type = iota
	Eval
)

// Source defines the interface for a source. Example sources are a source
// file or an evaluated string.
type Source interface {
	// Name returns the name of the source.
	Name() string

	// Type returns the type of the source.
	Type() Type

	// GetReader returns a new io.Reader for the source.
	GetReader() (io.Reader, error)
}

// Line defines the line in a source.
type Line = uint

// Column defines the column in a source line.
type Column = uint16

// Constants for unknown source positions.
const (
	LineUnknown   Line   = ^uint(0)
	ColumnUnknown Column = ^uint16(0)
)

// Cursor defines a precise position in a piece of source code.
type Cursor struct {
	Line Line
	Col  Column
}

// Span defines a range of characters in a piece of source code, possibly
// spanning multiple lines.
type Span struct {
	Source Source
	Beg    Cursor
	End    Cursor
}
