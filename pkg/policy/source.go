package policy

import (
	"context"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

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
	u, err := url.Parse(source)
	if err != nil {
		return nil, nil, err
	}
	pathParts := strings.Split(u.Path, "//")
	policyDir := filepath.Join(installDir, filepath.Base(source))
	if len(pathParts) > 1 {
		policyDir = filepath.Join(installDir, pathParts[1])
	}
	if err := getter.Get(ctx, policyDir, source); err != nil {
		return nil, nil, fmt.Errorf("failed to get source %s: %w", source, err)
	}

	data, err := afero.ReadFile(afero.NewOsFs(), filepath.Join(policyDir, defaultPolicyFileName))
	if err != nil {
		// TODO: make more descriptive error
		return nil, nil, fmt.Errorf("failed to open source: %w", err)
	}
	detectorType, _, err := getter.DetectType(source)
	if err != nil {
		return nil, nil, err
	}

	return data, &Meta{
		Type:      detectorType,
		Version:   u.Query().Get("ref"),
		SubPath:   u.Query().Get("subpolicy"),
		Directory: policyDir,
	}, nil
}
