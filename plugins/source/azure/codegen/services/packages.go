package services

import (
	"fmt"
	"path"
	"strings"

	"golang.org/x/exp/maps"
)

func packageName(x any) string {
	varType := strings.TrimPrefix(fmt.Sprintf("%T", x), "*")
	return strings.TrimSuffix(varType, path.Ext(varType))
}

func uniqPackages(packages ...string) []string {
	imports := make(map[string]struct{})
	for _, pkg := range packages {
		if len(pkg) > 0 {
			imports[pkg] = struct{}{}
		}
	}

	return maps.Keys(imports)
}
