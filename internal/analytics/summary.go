package analytics

import (
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
)

type DiagnosticsSummary struct {
	Total      int            `json:"total,omitempty"`
	ByType     map[string]int `json:"by_type,omitempty"`
	BySeverity map[string]int `json:"by_severity,omitempty"`
}

func SummarizeDiagnostics(diags diag.Diagnostics) DiagnosticsSummary {
	summary := DiagnosticsSummary{
		Total:      0,
		ByType:     make(map[string]int),
		BySeverity: make(map[string]int),
	}
	for _, d := range diags {
		summary.Total++
		severity := strings.ToLower(d.Severity().String())
		dtype := strings.ToLower(d.Type().String())

		if count, ok := summary.BySeverity[severity]; ok {
			summary.BySeverity[severity] = count + 1
		} else {
			summary.BySeverity[severity] = 1
		}

		if count, ok := summary.ByType[dtype]; ok {
			summary.ByType[dtype] = count + 1
		} else {
			summary.ByType[dtype] = 1
		}
	}
	return summary
}
