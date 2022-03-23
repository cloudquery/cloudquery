package telemetry

import (
	"errors"
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/stretchr/testify/assert"
)

func TestShouldIgnoreDiag(t *testing.T) {
	cases := []struct {
		D        diag.Diagnostic
		Expected bool
	}{
		{
			D:        diag.NewBaseError(errors.New("failed to connect to `host= user= database=`: server error (FATAL: the database system is in recovery mode (SQLSTATE 57P03))"), diag.DATABASE),
			Expected: true,
		},
		{
			D:        diag.NewBaseError(errors.New(`wafv2.managed_rule_groups: failed to insert to table "aws_wafv2_managed_rule_groups": ERROR: duplicate key value violates unique constraint "0_0_aws_wafv2_managed_rule_groups_cq_fetch_date_cq_id_key" (SQLSTATE 23505)`), diag.DATABASE),
			Expected: false,
		},
	}
	for i, tc := range cases {
		res := ShouldIgnoreDiag(tc.D)
		assert.Equalf(t, tc.Expected, res, "Test case #%d", i+1)
	}
}
