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

	t.Run("successful run without progress bar", func(t *testing.T) {
		filePath := filepath.Join(testDir, "cloudquery_linux.zip")
		if err := osFs.DownloadFile(ctx, filePath, testFileDownloadUrl, nil); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("successful run with progress bar", func(t *testing.T) {
		pu := console.NewProgress(ctx, func(o *console.ProgressOptions) {
			o.AppendDecorators = []decor.Decorator{decor.Percentage()}
		})
		progressCB := ui.CreateProgressUpdater(pu, "test-123")
		filePath := filepath.Join(testDir, "cloudquery_linux.zip")
		if err := osFs.DownloadFile(ctx, filePath, testFileDownloadUrl, progressCB); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("failure with non existing path", func(t *testing.T) {
		// we have to use real fs because mem fs creates non existing folders on the fly
		osFs.SetFSInstance(afero.NewOsFs())
		if err := osFs.DownloadFile(ctx, "/no/such/dir/cloudquery_linux.zip", testFileDownloadUrl, nil); err == nil {
			t.Fatal("expected to get non nil error")
		}
	})
}
