package main

import (
	"github.com/cloudquery/cloudquery/sdk"
	"github.com/cloudquery/cq-provider-azure/provider"
	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		JSONFormat: true,
	})

	p := &provider.Provider{Logger: logger.Named("azure")}

	sdk.ServePlugin(sdk.ServeOpts{
		Name:                "azure",
		Provider:            p,
		Logger:              logger,
		NoLogOutputOverride: false,
	})
}