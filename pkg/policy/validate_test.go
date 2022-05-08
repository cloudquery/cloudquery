package policy

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/internal/test"
	"github.com/cloudquery/cloudquery/pkg/core/database"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"github.com/stretchr/testify/assert"
)

type validateTest struct {
	Name          string
	Policy        *Policy
	ExpectedDiags []diag.FlatDiag
}

func TestValidate(t *testing.T) {

	testCases := []validateTest{
		{
			Name: "simple-validate",
			Policy: &Policy{
				Name:   "localPolicy",
				Source: "tests/validate/simple",
			},
			ExpectedDiags: nil,
		},
	}
	dsn := test.SetupDB(t)
	storage := database.NewStorage(dsn, nil)
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			diags := Validate(context.Background(), storage, &ValidateRequest{tc.Policy, "tests/output"})
			if tc.ExpectedDiags != nil {
				assert.ElementsMatch(t, tc.ExpectedDiags, diag.FlattenDiags(diags, false))
			} else {
				assert.Equal(t, []diag.FlatDiag{}, diag.FlattenDiags(diags, false))
			}
		})
	}
}
