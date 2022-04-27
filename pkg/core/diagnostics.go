package core

import "github.com/cloudquery/cq-provider-sdk/provider/diag"

type FetchDiagnostic struct {
	diag.Diagnostic
	Provider string
	Version  string
}

func convertToFetchDiags(diags diag.Diagnostics, provider, version string) diag.Diagnostics {
	fd := make(diag.Diagnostics, len(diags))
	for i, d := range diags {
		fd[i] = FetchDiagnostic{
			Diagnostic: d,
			Provider:   provider,
			Version:    version,
		}
	}
	return fd
}

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
		summary.Total += 1
		if _, ok := summary.BySeverity[d.Severity().String()]; ok {
			summary.BySeverity[d.Severity().String()] = 1
		} else {
			summary.BySeverity[d.Severity().String()] += 1
		}

		if _, ok := summary.ByType[d.Type().String()]; ok {
			summary.ByType[d.Type().String()] = 1
		} else {
			summary.ByType[d.Type().String()] += 1
		}
	}
	return summary
}
