package cmd_test

import (
	"strings"
	"testing"
	"time"

	"github.com/schollz/progressbar/v3"
)

func expectBuffer(t *testing.T, buf *strings.Builder, expect string) {
	t.Helper()
	current := strings.TrimSpace(buf.String())
	if current != expect {
		r := strings.NewReplacer("\r", "\\R", "\n", "\\N")
		current, expect = r.Replace(current), r.Replace(expect)
		t.Fatalf("Render mismatch\nResult: '%s'\nExpect: '%s'\n", current, expect)
	}
}

func TestRegression(t *testing.T) {
	buf := strings.Builder{}

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetWriter(&buf),
		progressbar.OptionSetDescription("Syncing resources..."),
		progressbar.OptionSetItsString("resources"),
		progressbar.OptionShowIts(),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionShowCount(),
	)
	bar.Reset()
	time.Sleep(1 * time.Second)
	expectBuffer(t, &buf, "")
	_ = bar.Add(5)
	expectBuffer(t, &buf, "- Syncing resources... (5/-, 5 resources/s) [1s]")
	time.Sleep(1 * time.Second)
	_ = bar.Add(5)
	expectBuffer(t, &buf, "- Syncing resources... (5/-, 5 resources/s) [1s] \r                                                 \r\r| Syncing resources... (10/-, 5 resources/s) [2s]")
	time.Sleep(1 * time.Second)
	_ = bar.Finish()
	expectBuffer(t, &buf, "- Syncing resources... (5/-, 5 resources/s) [1s] \r                                                 \r\r| Syncing resources... (10/-, 5 resources/s) [2s] \r                                                  \r\r- Syncing resources... (10/-, 3 resources/s) [3s]")
}

func TestRegressionClearOnFinish(t *testing.T) {
	buf := strings.Builder{}

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetWriter(&buf),
		progressbar.OptionSetDescription("Syncing resources..."),
		progressbar.OptionSetItsString("resources"),
		progressbar.OptionShowIts(),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionShowCount(),
		progressbar.OptionClearOnFinish(),
	)
	bar.Reset()
	time.Sleep(1 * time.Second)
	expectBuffer(t, &buf, "")
	_ = bar.Add(5)
	expectBuffer(t, &buf, "- Syncing resources... (5/-, 5 resources/s) [1s]")
	time.Sleep(1 * time.Second)
	_ = bar.Add(5)
	expectBuffer(t, &buf, "- Syncing resources... (5/-, 5 resources/s) [1s] \r                                                 \r\r| Syncing resources... (10/-, 5 resources/s) [2s]")
	time.Sleep(1 * time.Second)
	_ = bar.Finish()
	expectBuffer(t, &buf, "- Syncing resources... (5/-, 5 resources/s) [1s] \r                                                 \r\r| Syncing resources... (10/-, 5 resources/s) [2s]")
}
