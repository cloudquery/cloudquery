package ui

import (
	"io"

	"github.com/google/uuid"
)

const (
	StatusOK         = "ok"
	StatusError      = "error"
	StatusWarn       = "warn"
	StatusInProgress = "in_progress"
	StatusInfo       = "info"
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

	// MarkAllDone marks all progress bars as done
	MarkAllDone()
}

type ProgressUpdateFunc func(io.Reader, int64) io.Reader

// CreateProgressUpdater creates a progress update callback method for periodic updates.
func CreateProgressUpdater(progress Progress, displayName string) ProgressUpdateFunc {
	return func(reader io.Reader, total int64) io.Reader {
		id := uuid.New()
		progress.Add(id.String(), displayName, "downloading...", total+2)
		return progress.AttachReader(id.String(), reader)
	}
}
