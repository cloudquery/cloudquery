package azparser

import (
	"os"
	"strings"

	"golang.org/x/mod/modfile"
)

const (
	azureResourceManagerRootPackage = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager"
)

func GetArmModules(gomodPath string) ([]string, error) {
	var modules []string
	content, err := os.ReadFile(gomodPath)
	if err != nil {
		return nil, err
	}
	mod, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		return nil, err
	}

	for _, req := range mod.Require {
		if strings.HasPrefix(req.Mod.Path, azureResourceManagerRootPackage) && !packagesToSkip[req.Mod.Path] {
			modules = append(modules, req.Mod.String())
		}
	}

	return modules, nil
}
