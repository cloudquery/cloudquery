package ui

import (
	"io"
)

const (
	StatusOK         = "ok"
	StatusError      = "error"
	StatusWarn       = "warn"
	StatusTimeout    = "timeout"
	StatusAbort      = "abort"
	StatusInProgress = "in_progress"
)

// Progress is used to provide an updating progress to the user. The progress
// usually has a bar
type Progress interface {
	// Add adds an additional bar to the progress
	Add(id, displayName, message string, total int64)

	// Update writes a new status. This should be a single line.
	Update(id, status, msg string, amount int)

	// Increment the progress
	Increment(id string, amount int)

	AttachReader(name string, data io.Reader)

	Wait()
}
