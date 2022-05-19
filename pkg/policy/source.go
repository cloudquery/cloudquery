package policy

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cloudquery/cloudquery/internal/getter"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
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
			SubPolicy: subPolicy,
		},
	}, true, nil
}

// a list of regexps to match against error string to decide if it's a user error
var userErrors = []*regexp.Regexp{
	regexp.MustCompile("subdir .+? not found"),
	regexp.MustCompile("stat .*: permission denied"),
}

func classifyError(source string, err error) error {
	formattedError := fmt.Errorf("failed to get source %s: %w", source, err)
	errmsg := err.Error()
	for _, r := range userErrors {
		if r.MatchString(errmsg) {
			return diag.FromError(formattedError, diag.USER)
		}
	}
	return formattedError
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
		return nil, nil, classifyError(source, err)
	}

	data, err := afero.ReadFile(afero.NewOsFs(), filepath.Join(policyDir, defaultPolicyFileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil, diag.FromError(fmt.Errorf("could not find %q in %q. Please verify it exists", defaultPolicyFileName, source), diag.USER)
		}
		// TODO: make more descriptive error
		return nil, nil, fmt.Errorf("failed to open source: %w", err)
	}

	return data, &Meta{
		Type:      detectorType,
		Version:   version,
		SubPolicy: subPolicy,
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
