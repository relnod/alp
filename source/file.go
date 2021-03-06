package source

import (
	"io"
	"os"
)

// FileSource implements the source for a file.
type FileSource struct {
	filename string
}

// NewFileSource returns a new file source from a file name.
func NewFileSource(filename string) FileSource {
	return FileSource{filename: filename}
}

// Name returns the file name of the source file.
func (s *FileSource) Name() string {
	return s.filename
}

// Type returns source.File
func (s *FileSource) Type() Type {
	return File
}

// NewReader returns a new io.Reader, thst reads the contents of the source
// file.
func (s *FileSource) NewReader() (io.Reader, error) {
	f, err := os.Open(s.filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}
