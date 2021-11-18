package console

import (
	"testing"

	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/google/go-cmp/cmp"
)

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
			name: "no data or matching columns",
		},
		{
			name:    "no data",
			columns: []string{"arn"},
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
