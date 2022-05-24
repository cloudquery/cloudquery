package policy

import (
	"context"
	"sort"
	"testing"

	"github.com/cloudquery/cloudquery/internal/test"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/stretchr/testify/assert"
)

type validateTest struct {
	Name          string
	Policy        *Policy
	ExpectedDiags diag.FlatDiags
}

func TestValidate(t *testing.T) {
	testCases := []validateTest{
		{
			Name: "simple-validate",
			Policy: &Policy{
				Name:   "localPolicy",
				Source: "tests/validate/simple",
			},
		},
		{
			Name: "missing_identifiers",
			Policy: &Policy{
				Name:   "localPolicy",
				Source: "tests/validate/missing_identifiers",
			},
			ExpectedDiags: diag.FlatDiags{{Err: "check \"test_policy/1\" is missing identifier id", Type: diag.USER, Severity: diag.WARNING,
				Summary:     "check \"test_policy/1\" is missing identifier id",
				Description: diag.Description{Summary: "check \"test_policy/1\" is missing identifier id", Detail: "Check that query returns the following identifier: id"}}},
		},
		{
			Name: "identifiers_from_parent",
			Policy: &Policy{
				Name:   "localPolicy",
				Source: "tests/validate/identifiers_from_parent",
			},
		},
		{
			Name: "child_change_identifiers",
			Policy: &Policy{
				Name:   "localPolicy",
				Source: "tests/validate/child_change_identifiers",
			},
		},
		{
			Name: "missing_reason",
			Policy: &Policy{
				Name:   "localPolicy",
				Source: "tests/validate/no_reason",
			},
			ExpectedDiags: diag.FlatDiags{{Err: "check \"test_policy/1\" doesn't define reason in configuration or query",
				Type: diag.USER, Severity: diag.WARNING, Summary: "check \"test_policy/1\" doesn't define reason in configuration or query",
				Description: diag.Description{Summary: "check \"test_policy/1\" doesn't define reason in configuration or query",
					Detail: "Either add cq_reason column to query or reason attribute to check block"}}},
		},
	}
	dsn := test.SetupDB(t)
	storage := database.NewStorage(dsn, nil)
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			diags := Validate(context.Background(), storage, &ValidateRequest{tc.Policy, "tests/output"})
			if len(tc.ExpectedDiags) > 0 {
				sort.Sort(tc.ExpectedDiags)
				assert.ElementsMatch(t, tc.ExpectedDiags, diag.FlattenDiags(diags, false))
			} else {
				assert.Equal(t, diag.FlatDiags{}, diag.FlattenDiags(diags, false))
			}
		})
	}
}
