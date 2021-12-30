package policy

import (
	"context"
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/cloudquery/cloudquery/internal/getter"
	"github.com/spf13/afero"
)

const defaultPolicyFileName = "policy.hcl"

func DetectPolicy(name string, subPolicy string) (*Policy, bool, error) {
	t, found, err := getter.DetectType(name)

	if err != nil {
		return nil, false, fmt.Errorf("failed to detect policy in hub: %w", err)
	}
	if !found {
		return nil, false, nil
	}
	// TODO: parse version etc' from source
	return &Policy{
		Name:   name,
		Source: name,
		meta: &Meta{
			Type:    t,
			Version: "latest",
			SubPath: subPolicy,
		},
	}, true, nil
}

func LoadSource(ctx context.Context, installDir, source string) ([]byte, *Meta, error) {

	source, subPolicy := getter.SplitPackageSubDir(source)
	u, err := url.Parse(source)
	if err != nil {
		return nil, nil, err
	}
	detectorType, _, err := getter.DetectType(source)
	if err != nil {
		return nil, nil, err
	}
	policyDir := filepath.Join(installDir, getter.NormalizePath(source))
	if detectorType == "local" {
		policyDir = filepath.Join(installDir, filepath.Base(getter.NormalizePath(source)))
	}
	if err := getter.Get(ctx, policyDir, source); err != nil {
		return nil, nil, fmt.Errorf("failed to get source %s: %w", source, err)
	}

	data, err := afero.ReadFile(afero.NewOsFs(), filepath.Join(policyDir, defaultPolicyFileName))
	if err != nil {
		// TODO: make more descriptive error
		return nil, nil, fmt.Errorf("failed to open source: %w", err)
	}

	return data, &Meta{
		Type:      detectorType,
		Version:   u.Query().Get("ref"),
		SubPath:   subPolicy,
		Directory: policyDir,
	}, nil
}
