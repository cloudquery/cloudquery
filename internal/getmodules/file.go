package getmodules

import (
	"path/filepath"

	"github.com/hashicorp/go-getter"
	"github.com/spf13/afero"
)

// fileDetector is a replacement for go-getter's own file detector which also verifies the source file exists
type fileDetector struct {
	fd getter.FileDetector
}

func (d *fileDetector) Detect(src, pwd string) (string, bool, error) {
	if len(src) == 0 {
		return "", false, nil
	}
	checkPath := src
	if pwd != "" && !filepath.IsAbs(src) {
		checkPath = filepath.Join(pwd, src)
	}
	if ok, _ := afero.Exists(afero.NewOsFs(), checkPath); !ok {
		return "", false, nil
	}
	return d.fd.Detect(src, pwd)
}
