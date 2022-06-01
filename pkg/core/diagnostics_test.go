package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
)

func Test_SummarizeDiagnostics(t *testing.T) {

	testCases := []struct {
		Name            string
		Diags           diag.Diagnostics
		ExpectedSummary DiagnosticsSummary
	}{
		{
			Name:            "simple",
			Diags:           diag.FromError(errors.New("1"), diag.USER),
			ExpectedSummary: DiagnosticsSummary{Total: 1, ByType: map[string]int{"user": 1}, BySeverity: map[string]int{"error": 1}},
		},

		{
			Name: "multi-severity",
			Diags: diag.Diagnostics{}.Add(
				diag.FromError(errors.New("1"), diag.USER),
				diag.FromError(errors.New("1"), diag.USER, diag.WithSeverity(diag.WARNING)),
			),
			ExpectedSummary: DiagnosticsSummary{Total: 2, ByType: map[string]int{"user": 2}, BySeverity: map[string]int{"error": 1, "warning": 1}},
		},
		{
			Name: "multi-type",
			Diags: diag.Diagnostics{}.Add(
				diag.FromError(errors.New("1"), diag.USER),
				diag.FromError(errors.New("1"), diag.DATABASE),
				diag.FromError(errors.New("1"), diag.USER, diag.WithSeverity(diag.WARNING)),
			),
			ExpectedSummary: DiagnosticsSummary{Total: 3, ByType: map[string]int{"user": 2, "database": 1}, BySeverity: map[string]int{"error": 2, "warning": 1}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedSummary, SummarizeDiagnostics(tc.Diags))
		})
	}
}
