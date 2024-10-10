package publish

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickContentType(t *testing.T) {
	const dataDir = "testdata/content-type"

	tests := []struct {
		filename string
		want     string
		tryNoExt bool
	}{
		{
			filename: "png.png",
			want:     "image/png",
			tryNoExt: true,
		},
		{
			filename: "octetstream.svg",
			want:     "image/svg+xml",
		},
		{
			filename: "textplain.svg",
			want:     "image/svg+xml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			fullpath := filepath.Join(dataDir, tt.filename)
			qt, err := quickContentType(fullpath)
			require.NoError(t, err)
			require.Equal(t, tt.want, qt)

			if tt.tryNoExt {
				t.Run(tt.filename+"-noext", func(t *testing.T) {
					fullpath, err := copyTempFile(t, fullpath, "unknown-file")
					require.NoError(t, err)

					qt, err := quickContentType(fullpath)
					require.NoError(t, err)
					require.Equal(t, tt.want, qt)
				})
			}
		})
	}
}

func copyTempFile(t *testing.T, src, destName string) (string, error) {
	t.Helper()

	dir := t.TempDir()
	fullpath := filepath.Join(dir, destName)

	fp, err := os.Create(fullpath)
	require.NoError(t, err)
	defer fp.Close()

	srcHandle, err := os.Open(src)
	require.NoError(t, err)
	defer srcHandle.Close()

	_, err = io.Copy(fp, srcHandle)
	require.NoError(t, err)
	return fullpath, nil
}
