package testing

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/spf13/afero"
)

type OsFs struct {
	// fs is the afero filesystem instance
	fs afero.Fs
}

var (
	once         sync.Once
	osFsInstance *OsFs
)

func NewOsFs() *OsFs {
	// Singleton instantiation
	once.Do(func() {
		osFsInstance = &OsFs{
			fs: afero.NewOsFs(),
		}
	})
	return osFsInstance
}

func (o *OsFs) Stat(path string) (os.FileInfo, error) {
	return o.fs.Stat(path)
}

var defaultSupportedPolicyExtensions = []string{"hcl", "json"}
var defaultPolicyFileName = "policy"

func ParsePolicy(policyPath, policyDirectory string) (policy.Policies, error) {

	osFs := NewOsFs()
	// "/Users/benbernays/Documents/GitHub/aws"
	var policyFolder = policyDirectory
	// src := "/Users/benbernays/Documents/GitHub/aws/public_egress/policy.hcl"
	src := policyPath
	var policyFilePath = policyPath

	// check if abs path has provided if no, get the default policy file
	if !strings.HasSuffix(src, ".hcl") && !strings.HasSuffix(src, ".json") {
		policyFolder := src
		if info, err := osFs.Stat(policyFolder); err != nil || !info.IsDir() {
			log.Printf("could not find policy '%s' in the folder '%s'. Try to download the policy first", "unknown", src)
			os.Exit(1)
		}
		log.Println("internal repo folder set", "path", policyFolder)

		// Make sure policy file exists
		for _, extensionName := range defaultSupportedPolicyExtensions {
			currPolicyFile := filepath.Join(policyFolder, fmt.Sprintf("%s.%s", defaultPolicyFileName, extensionName))
			if _, err := osFs.Stat(currPolicyFile); err == nil {
				policyFilePath = currPolicyFile
				break
			}
		}
	}
	parser := config.NewParser()
	policiesRaw, diags := parser.LoadHCLFile(policyFilePath)
	if diags != nil && diags.HasErrors() {
		log.Printf("failed to load policy file: %#v", diags.Error())
	}

	policy, diagsDecode := policy.DecodePolicy(policiesRaw, diags, policyFolder)
	if diagsDecode != nil && diagsDecode.HasErrors() {
		log.Printf("failed to parse policy file: %#v", diagsDecode.Error())
		os.Exit(1)
	}
	return policy.Policies, nil
}
