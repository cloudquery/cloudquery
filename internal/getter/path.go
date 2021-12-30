package getter

import (
	"net/url"
	"path"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-getter"
)

func SplitPackageSubDir(given string) (packageAddr, subDir string) {
	packageAddr, subDir = getter.SourceDirSubdir(given)
	if subDir != "" {
		subDir = path.Clean(subDir)
	}
	return packageAddr, subDir
}

func NormalizePath(src string) string {
	srcParts := strings.Split(src, "::")
	src = srcParts[0]
	if len(srcParts) > 1 {
		src = srcParts[1]
	}
	_url, err := url.Parse(src)
	if err == nil {
		src = filepath.Join(_url.Host, _url.Path)
	}
	return filepath.ToSlash(strings.TrimRight(src, filepath.Ext(src)))
}
