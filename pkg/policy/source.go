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
	t, _, found, err := getter.DetectType(name)

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
			Type:      t,
			Version:   "latest",
			subPolicy: subPolicy,
		},
	}, true, nil
}

func LoadSource(ctx context.Context, installDir, source string) ([]byte, *Meta, error) {
	source, subPolicy := getter.ParseSourceSubPolicy(source)
	// parse syntactic URL holding @ instead of ?ref for params
	source, version := parseSyntacticUrl(source)
	if version == "" {
		u, _ := url.Parse(source)
		if u != nil {
			version = u.Query().Get("ref")
		}
	}

	detectorType, source, _, err := getter.DetectType(source)
	if err != nil {
		return nil, nil, err
	}
	policyDir := filepath.Join(installDir, getter.NormalizePath(source))
	if detectorType == "file" {
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
		Version:   version,
		subPolicy: subPolicy,
		Directory: policyDir,
	}, nil
}

func parseSyntacticUrl(source string) (string, string) {
	u := strings.Split(source, "@")
	if len(u) > 1 {
		return fmt.Sprintf("%s?ref=%s", u[0], u[1]), u[1]
	}
	return source, ""
}
