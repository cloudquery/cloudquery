package getter

import (
	"path"

	"github.com/hashicorp/go-getter"
)

func SplitPackageSubDir(given string) (packageAddr, subDir string) {
	packageAddr, subDir = getter.SourceDirSubdir(given)
	if subDir != "" {
		subDir = path.Clean(subDir)
	}
	return packageAddr, subDir
}
