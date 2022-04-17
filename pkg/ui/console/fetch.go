package console

import (
	"sort"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
)

func printFetchResponse(summary *core.FetchResponse, redactDiags, verbose bool) {
	if summary == nil {
		return
	}
	for _, pfs := range summary.ProviderFetchSummary {
		if len(pfs.Diagnostics()) > 0 {
			printDiagnostics("Fetch", pfs.Name, pfs.Diagnostics().Squash(), redactDiags, verbose)
			continue
		}
		ui.ColorizedOutput(ui.ColorWarning, "\n")
	}
}

func printDiagnostics(header, providerName string, diags diag.Diagnostics, redactDiags, verbose bool) {
	if redactDiags {
		diags = diags.Redacted()
	}

	if !verbose {
		var hasPrintableDiag bool
		for _, d := range diags {
			if d.Severity() != diag.IGNORE {
				hasPrintableDiag = true
				break
			}
		}
		if !hasPrintableDiag {
			return
		}
	}

	// sort diagnostics by severity/type
	sort.Sort(diags)

	ui.ColorizedOutput(ui.ColorHeader, "%s Diagnostics for provider %s:\n\n", header, providerName)
	for _, d := range diags {
		desc := d.Description()
		switch d.Severity() {
		case diag.IGNORE:
			if !verbose {
				continue
			}
			ui.ColorizedOutput(ui.ColorHeader, "Resource: %-10s Type: %-10s Severity: %s\n\tSummary: %s\n",
				ui.ColorProgress.Sprintf("%s", desc.Resource),
				ui.ColorProgressBold.Sprintf("%s", d.Type()),
				ui.ColorDebug.Sprintf("Ignore"),
				ui.ColorDebug.Sprintf("%s", desc.Summary))
		case diag.WARNING:
			ui.ColorizedOutput(ui.ColorHeader, "Resource: %-10s Type: %-10s Severity: %s\n\tSummary: %s\n",
				ui.ColorInfo.Sprintf("%s", desc.Resource),
				ui.ColorProgressBold.Sprintf("%s", d.Type()),
				ui.ColorWarning.Sprintf("Warning"),
				ui.ColorWarning.Sprintf("%s", desc.Summary))
		case diag.ERROR:
			ui.ColorizedOutput(ui.ColorHeader, "Resource: %-10s Type: %-10s Severity: %s\n\tSummary: %s\n",
				ui.ColorProgress.Sprintf("%s", desc.Resource),
				ui.ColorProgressBold.Sprintf("%s", d.Type()),
				ui.ColorErrorBold.Sprintf("Error"),
				ui.ColorErrorBold.Sprintf("%s", desc.Summary))
		}
		if len(desc.ResourceID) > 0 {
			ui.ColorizedOutput(ui.ColorInfo, "\tResource ID: %s\n", strings.Join(desc.ResourceID, ","))
		}
		if desc.Detail != "" {
			ui.ColorizedOutput(ui.ColorInfo, "\tDetail: %s\n", desc.Detail)
		}
	}
	ui.ColorizedOutput(ui.ColorInfo, "\n")
}
