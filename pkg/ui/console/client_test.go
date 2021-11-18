package console

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
