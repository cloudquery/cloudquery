package ui

import (
	"io"
)

const (
	StatusOK         = "ok"
	StatusError      = "error"
	StatusWarn       = "warn"
	StatusInProgress = "in_progress"
)

// Progress is used to provide an updating progress to the user. The progress
// usually has one or more bars
type Progress interface {
	// Add adds an additional bar to the progress
	Add(id, displayName, message string, total int64)

	// Update bars status, message and amount
	Update(id, status, msg string, amount int)

	// Increment the progress by given amount
	Increment(id string, amount int)

	// AttachReader to a progress so when an io is read the bar will update as well
	AttachReader(id string, data io.Reader) io.Reader

	// Wait for all progress bars to finish
	Wait()
}
