package terminal

import (
	"context"
	"io"
)

// UI is the primary interface for interacting with a user via the CLI.
//
// Some of the methods on this interface return values that have a lifetime
// such as Status and StepGroup. While these are still active (haven't called
// the close or equivalent method on these values), no other method on the
// UI should be called.
type UI interface {
	// Output outputs a message directly to the terminal. The remaining
	// arguments should be interpolations for the format string. After the
	// interpolations you may add Options.
	Output(string, ...interface{})


	// OutputWriters returns stdout and stderr writers. These are usually
	// but not always TTYs. This is useful for subprocesses, network requests,
	// etc. Note that writing to these is not thread-safe by default so
	// you must take care that there is only ever one writer.
	OutputWriters() (stdout, stderr io.Writer, err error)

	// Status returns a live-updating status that can be used for single-line
	// status updates that typically have a spinner or some similar style.
	// While a Status is live (Close isn't called), other methods on UI should
	// NOT be called.
	Status(ctx context.Context, name, message, total int64) Progress

	//// Table outputs the information formatted into a Table structure.
	//Table(*Table, ...Option)
	//
	//// StepGroup returns a value that can be used to output individual (possibly
	//// parallel) steps that have their own message, status indicator, spinner, and
	//// body. No other output mechanism (Output, Input, Status, etc.) may be
	//// called until the StepGroup is complete.
	//StepGroup() StepGroup
}