package console

import (
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/spf13/viper"
	"strconv"
)

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
	const (
		fetchSummary = "Provider %s fetch summary: %s Total Resources fetched: %d"
	)
	ui.ColorizedOutput(
		ui.ColorHeader,
		fetchSummary,
		key,
		s,
		summary.TotalResourcesFetched,
	)

	// errors
	errors := printIssues(diags, diag.ERROR, diag.PANIC)
	if len(errors) > 0 {
		const summaryErrors = "\t ❌ Errors: %s"
		ui.ColorizedOutput(ui.ColorHeader, summaryErrors, errors)
	}

	// warnings
	warnings := printIssues(diags, diag.WARNING)
	if len(warnings) > 0 {
		const summaryWarnings = "\t ⚠️ Warnings: %s"
		ui.ColorizedOutput(ui.ColorHeader, summaryWarnings, warnings)
	}

	// ignored issues
	ignored := printIssues(diags, diag.IGNORE)
	if len(ignored) > 0 {
		const summaryIgnored = "\t ❓ Ignored issues: %s"
		ui.ColorizedOutput(ui.ColorHeader, summaryIgnored, ignored)
		const footerIgnored = "\nProvider %s finished with %s ignored issues." +
			"\nThis may be normal, however, you can use `--verbose` flag to see more details"
		ui.ColorizedOutput(ui.ColorHeader, footerIgnored, key, ignored)
	}

	ui.ColorizedOutput(ui.ColorHeader, "\n\n")
}

func printIssues(diags diag.Diagnostics, severities ...diag.Severity) string {
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
		return
	}

	for _, sev := range sevs {
		deep += d.CountBySeverity(sev, true)
	}
	return
}
