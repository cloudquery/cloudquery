package console

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/cloudquery/cloudquery/pkg/ui"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
)

func printDiagnostics(header string, dd *diag.Diagnostics, redactDiags, verbose, squash bool) {
	// Nothing to
	if dd == nil || !dd.HasDiags() {
		return
	}
	diags := *dd

	if redactDiags {
		diags = diags.Redacted()
	}

	if squash {
		diags = diags.Squash()
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

	if header != "" {
		ui.ColorizedNoLogOutput(ui.ColorHeader, "%s Diagnostics:\n\n", header)
	} else {
		ui.ColorizedNoLogOutput(ui.ColorHeader, "Diagnostics:\n\n")
	}

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
		ui.ColorizedNoLogOutput(ui.ColorHeader, diagFormat, resourceInfo,
			ui.ColorProgressBold.Sprintf("%s", d.Type()),
			ui.ColorDebug.Sprintf("Ignore"),
			ui.ColorDebug.Sprintf("%s", desc.Summary))
		log.Debug().Stringer("type", d.Type()).Strs("resource", desc.ResourceID).Str("details", desc.Detail).Msg(desc.Summary)
	case diag.WARNING:
		ui.ColorizedNoLogOutput(ui.ColorHeader, diagFormat, resourceInfo,
			ui.ColorProgressBold.Sprintf("%s", d.Type()),
			ui.ColorWarning.Sprintf("Warning"),
			ui.ColorWarning.Sprintf("%s", desc.Summary))
		log.Warn().Stringer("type", d.Type()).Strs("resource", desc.ResourceID).Str("details", desc.Detail).Msg(desc.Summary)
	case diag.ERROR, diag.PANIC:
		ui.ColorizedNoLogOutput(ui.ColorHeader, diagFormat, resourceInfo,
			ui.ColorProgressBold.Sprintf("%s", d.Type()),
			ui.ColorErrorBold.Sprintf("Error"),
			ui.ColorErrorBold.Sprintf("%s", desc.Summary))
		log.Error().Stringer("type", d.Type()).Strs("resource", desc.ResourceID).Str("details", desc.Detail).Msg(desc.Summary)
	}
	if len(desc.ResourceID) > 0 {
		ui.ColorizedNoLogOutput(ui.ColorInfo, "\tResource ID: %s\n", strings.Join(desc.ResourceID, ","))
	}
	if desc.Detail != "" {
		ui.ColorizedNoLogOutput(ui.ColorInfo, "\tDetail: %s\n", desc.Detail)
	}
}
