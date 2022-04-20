package console

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/ui"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
)

func ClassifyDiagnostics(diags diag.Diagnostics) diag.Diagnostics {
	var classified diag.Diagnostics
	for _, d := range diags {
		if errors.Is(d, errors.New("invalid dsn")) {
			classified = classified.Add(diag.FromError(d, diag.DATABASE,
				diag.WithSummary("received dsn is invalid"),
				diag.WithDetails("received dsn is invalid, please check config/env it is correct")))
			continue
		}
		classified = classified.Add(d)
	}
	return classified

}

func printDiagnostics(_ string, diags diag.Diagnostics, redactDiags, verbose bool) {
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

	ui.ColorizedOutput(ui.ColorHeader, "Diagnostics:\n\n")
	for _, d := range diags {
		if !verbose && d.Severity() == diag.IGNORE {
			continue
		}
		printDiagnostic(d)
	}
	ui.ColorizedOutput(ui.ColorInfo, "\n")
}

const diagFormat = "%sType: %-10s Severity: %s\n\tSummary: %s\n"

func printDiagnostic(d diag.Diagnostic) {
	desc := d.Description()
	var resourceInfo = ""
	if desc.Resource != "" {
		resourceInfo = fmt.Sprintf("Resource: %-10s ", desc.Resource)
	}
	switch d.Severity() {
	case diag.IGNORE:
		ui.ColorizedOutput(ui.ColorHeader, diagFormat, resourceInfo,
			ui.ColorProgressBold.Sprintf("%s", d.Type()),
			ui.ColorDebug.Sprintf("Ignore"),
			ui.ColorDebug.Sprintf("%s", desc.Summary))
	case diag.WARNING:
		ui.ColorizedOutput(ui.ColorHeader, diagFormat, resourceInfo,
			ui.ColorProgressBold.Sprintf("%s", d.Type()),
			ui.ColorWarning.Sprintf("Warning"),
			ui.ColorWarning.Sprintf("%s", desc.Summary))
	case diag.ERROR:
		ui.ColorizedOutput(ui.ColorHeader, diagFormat, resourceInfo,
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
