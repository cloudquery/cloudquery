package getter

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/hashicorp/go-getter"
	"github.com/spf13/afero"
)

// fileDetector is a replacement for go-getter's own file detector which also verifies the source file exists
type fileDetector struct {
	fd               getter.FileDetector
	allowNonExisting bool
}

func isValidPath(fs afero.Fs, path string) bool {
	err := fs.MkdirAll(path, 0700)
	if err == nil {
		_ = fs.Remove(path)
	}

	return err == nil
}

func NewFileDetector(allowNonExisting bool) *fileDetector {
	p := new(fileDetector)
	p.allowNonExisting = allowNonExisting
	return p
}

// Extracted from getter.FileDetector
func fmtFileURL(path string) string {
	if runtime.GOOS == "windows" {
		// Make sure we're using "/" on Windows. URLs are "/"-based.
		path = filepath.ToSlash(path)
		return fmt.Sprintf("file://%s", path)
	}

	// Make sure that we don't start with "/" since we add that below.
	if path[0] == '/' {
		path = path[1:]
	}
	return fmt.Sprintf("file:///%s", path)
}

func (d *fileDetector) Detect(src, pwd string) (string, bool, error) {
	if len(src) == 0 {
		return "", false, nil
	}
	checkPath := src
	if pwd != "" && !filepath.IsAbs(src) {
		checkPath = filepath.Join(pwd, src)
	}
	fs := afero.NewOsFs()
	if ok, _ := afero.Exists(fs, checkPath); !ok {
		if d.allowNonExisting && isValidPath(fs, checkPath) {
			return fmtFileURL(checkPath), true, nil
		}
		return "", false, nil
	}
	return d.fd.Detect(src, pwd)
}
