package executor

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileReference struct {
	path     string
	filename string
	data     []byte
}

func (f FileReference) Filename() string {
	return f.filename
}

func (f FileReference) Data() (io.ReadCloser, int64, error) {
	if f.data != nil {
		return io.NopCloser(bytes.NewReader(f.data)), int64(len(f.data)), nil
	}
	file, err := os.Open(f.path)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, -1, fmt.Errorf("File '%s' not found", f.path)
	}
	if err != nil {
		return nil, -1, fmt.Errorf("Error reading file '%s': %v", f.path, err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		return nil, -1, fmt.Errorf("Error reading file size '%s': %v", f.path, err)
	}
	return file, fileStat.Size(), nil
}

func NewFileReference(path string) *FileReference {
	return &FileReference{
		filename: filepath.Base(path),
		path:     path,
	}
}

func NewFileReferenceData(filename string, data []byte) *FileReference {
	return &FileReference{
		filename: filename,
		data:     data,
	}
}
