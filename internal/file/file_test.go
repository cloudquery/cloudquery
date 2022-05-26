package file_test

import (
	"bytes"
	"context"
	"io"
	"path/filepath"
	"testing"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/internal/network"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/vbauerster/mpb/v6/decor"
)

type testBuffer struct {
	*bytes.Buffer
}

func (testBuffer) Close() error {
	return nil
}

func TestOsFs_DownloadFile(t *testing.T) {
	osFs := file.NewOsFs()
	osFs.SetFSInstance(afero.NewMemMapFs())
	testDir := "test/cloudquery"
	if err := osFs.MkdirAll(testDir, 0755); err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	downloader := func(data string) network.Downloader {
		return network.NewFakeHttpGetDownloader(
			testBuffer{bytes.NewBufferString(data)},
			int64(len(data)),
			nil)
	}
	t.Run("successful run without progress bar", func(t *testing.T) {
		filePath := filepath.Join(testDir, "cloudquery_linux.zip")
		if err := osFs.DownloadFile(downloader("data"), filePath, nil); err != nil {
			t.Fatal(err)
		}

		f, err := osFs.Open(filePath)
		assert.NoError(t, err)
		b, err := io.ReadAll(f)
		assert.NoError(t, err)
		assert.Equal(t, "data", string(b))
	})

	t.Run("successful run with progress bar", func(t *testing.T) {
		pu := console.NewProgress(ctx, func(o *console.ProgressOptions) {
			o.AppendDecorators = []decor.Decorator{decor.Percentage()}
		})
		progressCB := ui.CreateProgressUpdater(pu, "test-123")
		filePath := filepath.Join(testDir, "cloudquery_linux.zip")
		if err := osFs.DownloadFile(downloader("otherdata"), filePath, progressCB); err != nil {
			t.Fatal(err)
		}

		f, err := osFs.Open(filePath)
		assert.NoError(t, err)
		b, err := io.ReadAll(f)
		assert.NoError(t, err)
		assert.Equal(t, "otherdata", string(b))
	})

	t.Run("failure with non existing path", func(t *testing.T) {
		// we have to use real fs because mem fs creates non existing folders on the fly
		osFs.SetFSInstance(afero.NewOsFs())
		const filepath = "/no/such/dir/cloudquery_linux.zip"
		if err := osFs.DownloadFile(downloader("data"), filepath, nil); err == nil {
			t.Fatal("expected to get non nil error")
		}

		_, err := osFs.Stat(filepath)
		assert.Error(t, err)
	})
}
