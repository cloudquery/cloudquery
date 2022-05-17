package console

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/cloudquery/cloudquery/internal/analytics"
	"github.com/google/uuid"
)

func TestCreateClient(t *testing.T) {
	_ = analytics.Init()

	_, filename, _, _ := runtime.Caller(0)
	fixtures := filepath.Join(filepath.Dir(filename), "fixtures")

	tests := []struct {
		name       string
		configPath string
		wantErr    bool
	}{
		{
			name:       "valid",
			configPath: filepath.Join(fixtures, "config.yaml"),
			wantErr:    false,
		},
		{
			name:       "invalid config",
			configPath: filepath.Join(fixtures, "boom.yaml"),
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := CreateClient(context.Background(), tt.configPath, false, nil, uuid.Nil); (err != nil) != tt.wantErr {
				t.Errorf("CreateClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDescribePolicies(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	fixtures := filepath.Join(filepath.Dir(filename), "fixtures")

	var tests = []struct {
		name         string
		policySource string
		configPath   string
	}{
		{
			name:         "remote policy with config.hcl",
			policySource: "aws",
			configPath:   filepath.Join(fixtures, "config.yml"),
		},
		{
			name:         "local policy with config.hcl",
			policySource: fmt.Sprintf("file::%s", filepath.Join(fixtures, "example-policy")),
			configPath:   filepath.Join(fixtures, "config.yml"),
		},
		{
			name:         "remote policy without config.hcl",
			policySource: "aws",
		},
		{
			name:         "local policy without config.hcl",
			policySource: fmt.Sprintf("file::%s", filepath.Join(fixtures, "example-policy")),
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			c, err := CreateClient(ctx, "", true, nil, uuid.Nil)
			if err != nil {
				t.Errorf("Case: %d - CreateClient() error = %v", i, err)
				return
			}

			err = c.DescribePolicies(ctx, tt.policySource)
			if err != nil {
				t.Errorf("Case: %d - DescribePolicies() error = %v", i, err)
				return
			}
		})
	}
}
