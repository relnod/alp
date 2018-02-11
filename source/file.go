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

// GetReader returns a new io.Reader, thst reads the contents of the source
// file.
func (s *FileSource) GetReader() (io.Reader, error) {
	f, err := os.Open(s.filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}
