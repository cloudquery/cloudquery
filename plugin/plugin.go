package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const defaultOrganization = "cloudquery"

func getProviderPath(name string, version string) (string, error) {
	org := defaultOrganization
	split := strings.Split(name, "/")
	if len(split) == 2 {
		org = split[0]
		name = split[1]
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	extension := ""
	if runtime.GOOS == "windows" {
		extension = ".exe"
	}
	return filepath.Join(workingDir, ".cq", "providers", org, name, fmt.Sprintf("%s-%s-%s%s", version, runtime.GOOS, runtime.GOARCH, extension)), nil
}