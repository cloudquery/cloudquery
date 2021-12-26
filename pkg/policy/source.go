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

func LoadSource(ctx context.Context, installDir, source string) ([]byte, *Meta, error) {
	u, err := url.Parse(source)
	if err != nil {
		return nil, nil, err
	}
	dir, name := getter.SplitPackageSubDir(u.Path)
	if name == "" {
		name = filepath.Base(dir)
	}
	policyDir := filepath.Join(installDir, u.Path)
	if err := getter.Get(ctx, policyDir, source); err != nil {
		return nil, nil, fmt.Errorf("failed to get source %s: %w", source, err)
	}

	data, err := afero.ReadFile(afero.NewOsFs(), filepath.Join(policyDir, defaultPolicyFileName))
	if err != nil {
		// TODO: make more descriptive error
		return nil, nil, fmt.Errorf("failed to open source: %w", err)
	}
	detectorType, err := getter.DetectType(source)
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
