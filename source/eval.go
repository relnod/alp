package source

import (
	"bytes"
	"io"
)

// EvalSource implements the source for an evaluated string
type EvalSource struct {
	name   string
	source string
}

// NewEvalSource returns a new eval source.
func NewEvalSource(name string, source string) EvalSource {
	return EvalSource{name: name, source: source}
}

// Name returns the name of the source.
func (s *EvalSource) Name() string {
	return s.name
}

// GetReader returns a new reader for the eval source. The Reader will return
// the whole source at once.
func (s *EvalSource) GetReader() (io.Reader, error) {
	return bytes.NewReader([]byte(s.source)), nil
}
