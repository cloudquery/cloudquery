package terminal

import (
	"context"
	"io"
	"os"
)

type ConsoleUI struct{}

func (c ConsoleUI) Output(str string, i ...interface{}) {
	return
}

func (c ConsoleUI) OutputWriters() (stdout, stderr io.Writer, err error) {
	return os.Stdout, os.Stderr, nil
}


func (c ConsoleUI) Progress(ctx context.Context, name, message string, total int64) Progress {
	return nil
}


