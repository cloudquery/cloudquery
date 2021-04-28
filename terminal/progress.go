package terminal

import (
	"context"
	"github.com/vbauerster/mpb/v6"
	"github.com/vbauerster/mpb/v6/decor"
)

const (
	StatusOK      = "ok"
	StatusError   = "error"
	StatusWarn    = "warn"
	StatusTimeout = "timeout"
	StatusAbort   = "abort"
)

var emojiStatus = map[string]string{
	StatusOK:      "✓",
	StatusError:   "❌",
	StatusWarn:    "⚠️",
	StatusTimeout: "⌛",
}

var textStatus = map[string]string{
	StatusOK:      " +",
	StatusError:   " !",
	StatusWarn:    " *",
	StatusTimeout: "<>",
}

// Progress is used to provide an updating progress to the user. The progress
// usually has a bar
type Progress interface {

	// Update writes a new status. This should be a single line.
	Update(msg string)

	// Increment the progress
	Increment(amount int)

	// Step Indicates that a step has finished, confering an ok, error, or warn upon
	// it's finishing state. If the status is not StatusOK, StatusError, or StatusWarn
	// then the status text is written directly to the output, allowing for custom
	// statuses.
	Step(status, msg string)

	// Close should be called when the live updating is complete. The
	// status will be cleared from the line.
	Close() error
}

type UiProgress struct {
	p *mpb.Progress
	b *mpb.Bar
	// name to display in progress
	name string
	// message to display in progress
	message string
	// status of progress
	status string
}

func NewUiProgress(ctx context.Context, name, message string, total int64) UiProgress {

	u := UiProgress{
		p:       mpb.NewWithContext(ctx),
		b:       nil,
		name:    name,
		message: message,
		status:  "",
	}
	bar := u.p.Add(total,
		// progress bar filler with customized style
		mpb.NewBarFiller("[=>-|"),
		mpb.BarFillerOnComplete(u.status),
		mpb.PrependDecorators(
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			decor.Any(func(statistics decor.Statistics) string {
				return u.message
			}, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "",
			),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)
	u.b = bar
	return u
}

func (u UiProgress) Increment(n int) {
	u.b.IncrBy(n)
}

func (u UiProgress) Update(msg string) {
	panic("implement me")
}

func (u UiProgress) Step(status, msg string) {
	u.message = msg
	u.status = status
}

func (u UiProgress) Close() error {
	return nil
}
