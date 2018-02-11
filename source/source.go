package source

import (
	"io"
)

// Source defines the interface for a source. Example sources are a source
// file or an evaluated string.
//
// TODO: SourceMap
//
type Source interface {
	// Name returns the name of the source
	Name() string

	// GetReader returns a new io.Reader for the source
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
