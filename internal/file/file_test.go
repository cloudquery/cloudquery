package file_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/afero"
	"github.com/vbauerster/mpb/v6/decor"
)

const testFileDownloadUrl = "https://github.com/cloudquery/cloudquery/releases/download/v0.13.5/cloudquery_Linux_arm64.zip"

func TestOsFs_DownloadFile(t *testing.T) {
	osFs := file.NewOsFs()
	osFs.SetFSInstance(afero.NewMemMapFs())
	testDir := "test/cloudquery"
	if err := osFs.MkdirAll(testDir, 0755); err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	filePath := filepath.Join(testDir, "cloudquery_linux.zip")
	if err := osFs.DownloadFile(ctx, filePath, testFileDownloadUrl, nil); err != nil {
		t.Fatal(err)
	}

	// With progress updater
	pu := console.NewProgress(ctx, func(o *console.ProgressOptions) {
		o.AppendDecorators = []decor.Decorator{decor.Percentage()}
	})
	progressCB := ui.CreateProgressUpdater(pu, "test-123")
	if err := osFs.DownloadFile(ctx, filePath, testFileDownloadUrl, progressCB); err != nil {
		t.Fatal(err)
	}
}
