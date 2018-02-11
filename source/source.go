// TODO: SourceMap

package source

// ID is the id of a source, i.e. source file or evaluated string.
type ID = uint

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
	Source ID
	Beg    Cursor
	End    Cursor
}
