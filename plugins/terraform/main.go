package main

import (
	"github.com/cloudquery/cloudquery/sdk"
	"github.com/cloudquery/cq-provider-template/resources"
)



func main() {
	sdk.ServePlugin(sdk.ServeOpts{
		Name:                "aws",
		Provider:            resources.Provider(),
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
