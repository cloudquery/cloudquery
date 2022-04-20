package console

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/ui"
)

func printFetchResponse(summary *core.FetchResponse, redactDiags, verbose bool) {
	if summary == nil {
		return
	}
	for _, pfs := range summary.ProviderFetchSummary {
		if len(pfs.Diagnostics()) > 0 {
			printDiagnostics(fmt.Sprintf("fetch - %s", pfs.Name), pfs.Diagnostics().Squash(), redactDiags, verbose)
			continue
		}
		ui.ColorizedOutput(ui.ColorInfo, "\n")
	}
}
