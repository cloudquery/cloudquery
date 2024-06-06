package cmd

import "io"

type progressBar interface {
	io.Writer
	Add(num int) error
	Finish() error
}

type noopProgressBar struct {
}

func (noopProgressBar) Write(p []byte) (int, error) {
	return len(p), nil
}
func (noopProgressBar) Add(_ int) error {
	return nil
}
func (noopProgressBar) Finish() error {
	return nil
}
