package main

import (
	"github.com/cloudquery/cloudquery/sdk"
	"github.com/cloudquery/cq-provider-aws/resources"
)

func main() {
	sdk.ServePlugin(sdk.ServeOpts{
		Name:                "aws",
		Provider:            resources.Provider(),
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
