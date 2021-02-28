package main

import (
	"github.com/cloudquery/cloudquery/sdk"
	"github.com/cloudquery/cq-provider-gcp/provider"
	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		JSONFormat: true,
	})

	p := &provider.Provider{Logger: logger.Named("gcp")}

	sdk.ServePlugin(sdk.ServeOpts{
		Name:                "gcp",
		Provider:            p,
		Logger:              logger,
		NoLogOutputOverride: false,
	})
}