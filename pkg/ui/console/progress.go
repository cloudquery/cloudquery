package console

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/fatih/color"
	"github.com/vbauerster/mpb/v6"
	"github.com/vbauerster/mpb/v6/decor"
)

var emojiStatus = map[string]string{
	ui.StatusOK:         color.GreenString("✓"),
	ui.StatusInfo:       "📋",
	ui.StatusError:      color.RedString("❌"),
	ui.StatusWarn:       "⚠️",
	ui.StatusInProgress: "⌛",
}

type Bar struct {
	b       *mpb.Bar
	Name    string
	Message string
	Status  string
	Total   int64
}

func (b *Bar) SetTotal(total int64, triggerComplete bool) {
	b.Total = total
	b.b.SetTotal(total, triggerComplete)
}

func (b *Bar) Done() {
	b.b.Abort(false)
}

type Progress struct {
	p       *mpb.Progress
	bars    map[string]*Bar
	lock    sync.Mutex
	options *ProgressOptions
}

type ProgressOptions struct {
	Filler           string
	StatusFunc       func(b *Bar, statistics decor.Statistics) string
	MessageHook      func(b *Bar, statistics decor.Statistics) string
	AppendDecorators []decor.Decorator
}

type ProgressOption func(o *ProgressOptions)

func NewProgress(ctx context.Context, opts ...ProgressOption) *Progress {
	u := &Progress{
		p:    mpb.NewWithContext(ctx, mpb.WithWidth(64), mpb.WithRefreshRate(180*time.Millisecond)),
		bars: make(map[string]*Bar),
		options: &ProgressOptions{
			Filler:           "[=>-|",
			StatusFunc:       defaultStatusUpdater,
			MessageHook:      defaultMessageUpdater,
			AppendDecorators: nil,
		},
	}
	for _, o := range opts {
		o(u.options)
	}
	return u
}

func (u *Progress) Add(name, displayName, message string, total int64) {
	bar := u.p.Add(total, // total of file + 2 verify results
		// progress bar filler with customized style
		mpb.NewBarFiller(u.options.Filler),
		mpb.BarFillerOnComplete(""),
		mpb.PrependDecorators(
			decor.Any(
				func(statistics decor.Statistics) string {
					if u.options.StatusFunc == nil {
						return ""
					}
					u.lock.Lock()
					defer u.lock.Unlock()
					uiBar := u.bars[name]
					return u.options.StatusFunc(uiBar, statistics)
				}, decor.WC{W: 2, C: 1}),
			// display our name with one space on the right
			decor.Name(displayName, decor.WC{W: len(displayName) + 1, C: decor.DidentRight}),
			decor.Any(func(statistics decor.Statistics) string {
				if u.options.MessageHook == nil {
					return ""
				}
				u.lock.Lock()
				defer u.lock.Unlock()
				uiBar := u.bars[name]
				return u.options.MessageHook(uiBar, statistics)
			}, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			// replace ETA decorator with nothing, OnComplete event
			decor.Elapsed(decor.ET_STYLE_GO, decor.WC{W: 6}),
		),
		mpb.AppendDecorators(u.options.AppendDecorators...),
	)
	u.bars[name] = &Bar{
		b:       bar,
		Name:    name,
		Message: message,
		Status:  ui.StatusInProgress,
		Total:   total,
	}
}

func (u *Progress) Increment(name string, n int) {
	u.lock.Lock()
	defer u.lock.Lock()
	bar, ok := u.bars[name]
	if !ok {
		return
	}
	bar.b.IncrBy(n)
}

func (u *Progress) Update(name, status, msg string, n int) {
	u.lock.Lock()
	defer u.lock.Unlock()
	bar, ok := u.bars[name]
	if !ok {
		return
	}
	bar.Message = msg
	bar.Status = status
	if n > 0 {
		bar.b.IncrBy(n)
	}
}

func (u *Progress) AttachReader(name string, data io.Reader) io.Reader {
	u.lock.Lock()
	defer u.lock.Unlock()
	bar, ok := u.bars[name]
	if !ok {
		return data
	}
	return bar.b.ProxyReader(io.LimitReader(data, bar.Total))
}

func (u *Progress) Wait() {
	time.Sleep(600 * time.Millisecond)
	u.p.Wait()
	fmt.Println()
}

func (u *Progress) GetBar(name string) *Bar {
	u.lock.Lock()
	defer u.lock.Unlock()
	bar, ok := u.bars[name]
	if !ok {
		return nil
	}
	return bar
}

func (u *Progress) AbortAll() {
	u.lock.Lock()
	defer u.lock.Unlock()
	for _, b := range u.bars {
		b.b.Abort(true)
	}
	time.Sleep(100 * time.Millisecond)
	u.p.Wait()
	fmt.Println()
}

func (u *Progress) MarkAllDone() {
	u.lock.Lock()
	defer u.lock.Unlock()
	for _, b := range u.bars {
		if b.Status == ui.StatusInProgress {
			b.Status = ui.StatusOK
		}
		b.SetTotal(b.Total, true)
	}
}

func defaultStatusUpdater(u *Bar, _ decor.Statistics) string {
	return emojiStatus[u.Status]
}

func defaultMessageUpdater(u *Bar, _ decor.Statistics) string {
	return u.Message
}
