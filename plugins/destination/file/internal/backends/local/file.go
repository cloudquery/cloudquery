package local

import (
	"io"
	"io/fs"
	"os"

	"github.com/google/uuid"
)

type file struct {
	file *os.File
	written  uint64
	name string
}

func (f *file) Write(data []byte) (int, error) {
	n, err := f.file.Write(data)
	if err != nil {
		return n, err
	}
	
	f.written += uint64(n)
	if f.written >= 10000 {
		if err := f.file.Close(); err != nil {
			f.file = nil
			return n, err
		}
		f.written = 0
		name := uuid.NewString() + "." + f.name
		f.file, err = os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func (f *file) Read(p []byte) (n int, err error) {
	return f.file.Read(p)
}

func (f *file) Close() error {
	defer func() {
		f.file = nil
		f.written = 0
	}()
	if f.file == nil {
		return fs.ErrClosed
	}
	return f.file.Close()
}

func OpenAppendOnly(name string) (io.WriteCloser, error) {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	return &file{
		file: f,
		name: name,
	}, nil
}

func OpenReadOnly(name string) (io.ReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return &file{
		file: f,
		name: name,
	}, nil
}