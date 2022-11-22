package client

import (
	"fmt"
	"os"
)

type openMode int

const (
	// OpenModeAppend appends to the file
	openModeAppend openMode = iota
	// OpenModeReadOnly opens the file in read only mode
	openModeReadOnly
)

// generic file implements writer and reader interface
// for local, gcs, s3 storages
type genericFile struct {
	backend BackendType
	f 		 *os.File
	mode  openMode
}

func OpenReadOnly(name string, backend BackendType) (*genericFile, error) {
	var f *os.File
	var err error
	// io.Writer
	if backend == BackendTypeLocal {
		f, err = os.Open(name)
		if err != nil {
			return nil, err
		}
	}
	return &genericFile{
		backend: backend,
		f: f,
		mode: openModeReadOnly,
	}, nil
}

func OpenAppendOnly(name string, backend BackendType) (*genericFile, error) {
	var f *os.File
	var err error
	if backend == BackendTypeLocal {
		f, err = os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return nil, err
		}
	}
	return &genericFile{
		backend: backend,
		f: f,
		mode: openModeAppend,
	}, nil
}


func (f *genericFile) Write(data []byte) (n int, err error) {
	if f.mode != openModeAppend {
		return 0, fmt.Errorf("file is not opened in append mode")
	}
	switch f.backend {
	case BackendTypeLocal:
		return f.f.Write(data)
	case BackendTypeGCS:
		return 0, fmt.Errorf("not implemented")
	case BackendTypeS3:
		return 0, fmt.Errorf("not implemented")
	default:
		panic("unknown backend " + f.backend)
	}
}