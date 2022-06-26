package console

import (
	"strconv"

	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/spf13/viper"
)

// PrintProviderSummary is a helper to print the fetch summary in an easily readable format.
func PrintProviderSummary(summary *core.ProviderFetchSummary) {
	s := emojiStatus[ui.StatusOK]
	if summary.Status == core.FetchCanceled {
		s = emojiStatus[ui.StatusError] + " (canceled)"
	}
	key := summary.Name
	if summary.Name != summary.Alias {
		key = summary.Name + `(` + summary.Alias + `)`
	}
	diags := summary.Diagnostics().Squash()
	ui.ColorizedOutput(
		ui.ColorHeader,
		"Provider %s fetch summary: %s Total Resources fetched: %d",
		key,
		s,
		summary.TotalResourcesFetched,
	)

	// errors
	errors := formatIssues(diags, diag.ERROR, diag.PANIC)
	if len(errors) > 0 {
		ui.ColorizedOutput(ui.ColorHeader, "\t ❌ Errors: %s", errors)
	}

	// warnings
	warnings := formatIssues(diags, diag.WARNING)
	if len(warnings) > 0 {
		ui.ColorizedOutput(ui.ColorHeader, "\t ⚠️ Warnings: %s", warnings)
	}

	// ignored issues
	ignored := formatIssues(diags, diag.IGNORE)
	if len(ignored) > 0 {
		ui.ColorizedOutput(ui.ColorHeader, "\t ❓ Ignored issues: %s", ignored)
		ui.ColorizedOutput(ui.ColorHeader,
			"\nProvider %s finished with %s ignored issues."+
				"\nThis may be normal, however, you can use `--verbose` flag to see more details.",
			key, ignored)
	}

	ui.ColorizedOutput(ui.ColorHeader, "\n\n")
}

// formatIssues will pretty-print the diagnostics by the requested severities:
// - for no issues "" is returned
// - for any deep issues the "base (deep)" amounts are printed
// - for basic case with no deep issues but rather the base ones, the "base" amount is printed
func formatIssues(diags diag.Diagnostics, severities ...diag.Severity) string {
	basic, deep := countSeverity(diags, severities...)
	switch {
	case deep > 0:
		return strconv.FormatUint(basic, 10) + `(` + strconv.FormatUint(deep, 10) + `)`
	case basic > 0:
		return strconv.FormatUint(basic, 10)
	default:
		return ``
	}
}

func countSeverity(d diag.Diagnostics, sevs ...diag.Severity) (basic, deep uint64) {
	for _, sev := range sevs {
		basic += d.CountBySeverity(sev, false)
	}

	if !viper.GetBool("verbose") {
		return basic, 0
	}

	for _, sev := range sevs {
		deep += d.CountBySeverity(sev, true)
	}
	return basic, deep
}
