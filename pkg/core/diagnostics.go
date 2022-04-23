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
	Total      int                         `json:"total,omitempty"`
	ByType     map[diag.DiagnosticType]int `json:"by_type,omitempty"`
	BySeverity map[diag.Severity]int       `json:"by_severity,omitempty"`
}

// TODO: convert to map[string]int

func SummarizeDiagnostics(diags diag.Diagnostics) DiagnosticsSummary {
	summary := DiagnosticsSummary{
		Total:      0,
		ByType:     make(map[diag.DiagnosticType]int),
		BySeverity: make(map[diag.Severity]int),
	}
	for _, d := range diags {
		summary.Total += 1
		if _, ok := summary.BySeverity[d.Severity()]; ok {
			summary.BySeverity[d.Severity()] = 1
		} else {
			summary.BySeverity[d.Severity()] += 1
		}

		if _, ok := summary.ByType[d.Type()]; ok {
			summary.ByType[d.Type()] = 1
		} else {
			summary.ByType[d.Type()] += 1
		}
	}
	return summary
}
