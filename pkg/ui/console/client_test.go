package console

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/cloudquery/cloudquery/pkg/policy"
)

func TestCreateClient(t *testing.T) {
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
			if _, err := CreateClient(context.Background(), tt.configPath, nil); (err != nil) != tt.wantErr {
				t.Errorf("CreateClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDefineResultColumnWidths(t *testing.T) {
	var tests = []struct {
		name   string
		data   []*policy.QueryResult
		output string
	}{
		{
			name:   "no description or name",
			output: "\t%s  %-0s %-0s %10s",
		}, {
			name: "Only Description",
			data: []*policy.QueryResult{
				{Description: "test"},
			},
			output: "\t%s  %-0s %-5s %10s",
		},
		{
			name: "Only Name",
			data: []*policy.QueryResult{
				{Name: "test"},
			},
			output: "\t%s  %-5s %-0s %10s",
		},
		{
			name: "Multiple Names",
			data: []*policy.QueryResult{
				{Name: "test"},
				{Name: "test-test-test"},
			},
			output: "\t%s  %-15s %-0s %10s",
		},
		{
			name: "Multiple Descriptions",
			data: []*policy.QueryResult{
				{Description: "test"},
				{Description: "test-test-test"},
			},
			output: "\t%s  %-0s %-15s %10s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := defineResultColumnWidths(tt.data)

			diff := cmp.Diff(ans, tt.output)
			if diff != "" {
				t.Fatalf("values are not the same %s", diff)

			}
		})
	}

}

func TestFindOutput(t *testing.T) {
	var tests = []struct {
		name    string
		data    [][]interface{}
		columns []string
		output  []string
	}{
		{
			name:   "no data or matching columns",
			output: []string{},
		},
		{
			name:    "no data",
			columns: []string{"arn"},
			output:  []string{},
		},
		{
			name: "matching data and columns",
			data: [][]interface{}{
				{0, 1, 2, 3},
				{4, 5, 6, 7},
				{8, 9, 10, 11},
			},
			columns: []string{"arn"},
			output:  []string{"0", "4", "8"},
		},
		{
			name: "matching data and multiple columns",
			data: [][]interface{}{
				{0, 1, 2, 3},
				{4, 5, 6, 7},
				{8, 9, 10, 11},
			},
			columns: []string{"arn", "id", "uid", "uuid"},
			output:  []string{"1", "5", "9"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := findOutput(tt.columns, tt.data)

			diff := cmp.Diff(ans, tt.output)
			if diff != "" {
				t.Fatalf("values are not the same %s", diff)

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
			c, err := CreateNullClient(ctx)
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
