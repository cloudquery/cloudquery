package main

import (
	"github.com/cloudquery/cloudquery/sdk"
	"github.com/cloudquery/cq-provider-okta/provider"
	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		JSONFormat: true,
	})

	p := &provider.Provider{Logger: logger.Named("okta")}

	sdk.ServePlugin(sdk.ServeOpts{
		Name:                "okta",
		Provider:            p,
		Logger:              logger,
		NoLogOutputOverride: false,
	})
}
