package console

import (
	"context"
	"fmt"
	registry2 "github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/fatih/color"
	"github.com/vbauerster/mpb/v6"
	"github.com/vbauerster/mpb/v6/decor"
	"io"
	"sync"
	"time"
)

const (
	StatusOK         = "ok"
	StatusError      = "error"
	StatusWarn       = "warn"
	StatusTimeout    = "timeout"
	StatusAbort      = "abort"
	StatusInProgress = "in_progress"
)

//
var emojiStatus = map[string]string{
	StatusOK:         color.GreenString("✓"),
	StatusError:      color.RedString("❌"),
	StatusWarn:       "⚠️",
	StatusInProgress: "⌛",
}

// Progress is used to provide an updating progress to the user. The progress
// usually has a bar
type Progress interface {

	// Add adds an aditional bar to the progress
	Add(name string, message string, total int64)

	// Update writes a new status. This should be a single line.
	Update(msg string)

	// Increment the progress
	Increment(amount int)

	// Step Indicates that a step has finished, confering an ok, error, or warn upon
	// it's finishing state. If the status is not StatusOK, StatusError, or StatusWarn
	// then the status text is written directly to the output, allowing for custom
	// statuses.
	Step(status, msg string)
}

type UiBar struct {
	b       *mpb.Bar
	name    string
	message string
	status  string
}

type UiProgress struct {
	p       *mpb.Progress
	bars    map[string]*UiBar
	lock    sync.RWMutex
	options ProgressOptions
}

type ProgressOptions struct {
	filler           string
	statusFunc       func(b *UiBar, statistics decor.Statistics) string
	messageHook      func(b *UiBar, statistics decor.Statistics) string
	appendDecorators []decor.Decorator
}

func DefaultStatusUpdater(u *UiBar, s decor.Statistics) string {
	return emojiStatus[u.status]
}

func DefaultMessageUpdater(u *UiBar, s decor.Statistics) string {
	return u.message
}

// TODO: make this functional
func NewUiProgress(ctx context.Context, opts ProgressOptions) *UiProgress {
	u := &UiProgress{
		p:       mpb.NewWithContext(ctx, mpb.WithWidth(64), mpb.WithRefreshRate(180*time.Millisecond)),
		bars:    make(map[string]*UiBar),
		options: opts,
	}
	return u
}

func (u *UiProgress) Add(name, displayName, message string, total int64) {
	bar := u.p.Add(total, // total of file + 2 verify results
		// progress bar filler with customized style
		mpb.NewBarFiller("[=>-|"),
		mpb.BarFillerOnComplete(""),
		mpb.PrependDecorators(
			decor.Any(
				func(statistics decor.Statistics) string {
					if u.options.statusFunc == nil {
						return ""
					}
					u.lock.RLock()
					defer u.lock.RUnlock()
					uiBar := u.bars[name]
					return u.options.statusFunc(uiBar, statistics)
				}, decor.WC{W: 2, C: 1}),
			// display our name with one space on the right
			decor.Name(displayName, decor.WC{W: len(displayName) + 1, C: decor.DidentRight}),
			decor.Any(func(statistics decor.Statistics) string {
				if u.options.messageHook == nil {
					return ""
				}
				u.lock.RLock()
				defer u.lock.RUnlock()
				uiBar := u.bars[name]
				return u.options.messageHook(uiBar, statistics)
			}, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			// replace ETA decorator with nothing, OnComplete event
			decor.Elapsed(decor.ET_STYLE_GO, decor.WC{W: 4}),
		),
		mpb.AppendDecorators(u.options.appendDecorators...),
	)
	u.bars[name] = &UiBar{
		b:       bar,
		name:    name,
		message: message,
		status:  StatusInProgress,
	}
}

func (u *UiProgress) Increment(name string, n int) {
	u.lock.RLock()
	defer u.lock.RUnlock()
	bar, ok := u.bars[name]
	if !ok {
		return
	}
	bar.b.IncrBy(n)
}

func (u *UiProgress) Update(name, msg string) {
	u.lock.RLock()
	defer u.lock.RUnlock()
	bar, ok := u.bars[name]
	if !ok {
		return
	}
	bar.message = msg
}

func (u *UiProgress) Step(name, status, msg string) {
	u.lock.RLock()
	defer u.lock.RUnlock()
	bar, ok := u.bars[name]
	if !ok {
		return
	}
	bar.message = msg
	bar.status = status
}

func (u *UiProgress) Wait() {
	time.Sleep(100 * time.Millisecond)
	u.p.Wait()
}

type ProgressUpdater struct {
	progress  *mpb.Progress
	barStatus map[string]string
	bars      map[string]*mpb.Bar
	status    string
}

func (p ProgressUpdater) OnDownload(providerName string, version string, size int64, data io.Reader) io.Reader {
	total := size
	name := fmt.Sprintf("cq-provider-%s@%s", providerName, version)
	// adding a single bar, which will inherit container's width
	bar := p.progress.Add(total+2, // total of file + 2 verify results
		// progress bar filler with customized style
		mpb.NewBarFiller("[=>-|"),
		mpb.BarFillerOnComplete(p.status),
		mpb.PrependDecorators(
			decor.Any(func(statistics decor.Statistics) string {
				status := p.barStatus[providerName]
				if status == string(registry2.Verified) {
					return color.GreenString("✓")
				}
				if status == "downloading" || status == "verifying" {
					return color.WhiteString("⌛")
				}
				return color.RedString("❌")
			}, decor.WC{W: 2, C: 1}),
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			decor.Any(func(statistics decor.Statistics) string {
				if p.barStatus[providerName] == string(registry2.Verified) {
					return string(registry2.Verified)
				}
				return p.barStatus[providerName] + "..."
			}, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "",
			),
		),
		mpb.AppendDecorators(
			decor.Percentage(),
		),
	)
	p.barStatus[providerName] = "downloading"
	p.bars[providerName] = bar
	return bar.ProxyReader(io.LimitReader(data, size))
}

func (p ProgressUpdater) OnVerify(providerName string, status registry2.VerifyStatus) {
	p.barStatus[providerName] = string(status)
	p.bars[providerName].Increment()
}
