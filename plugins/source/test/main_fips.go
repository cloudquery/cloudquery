//go:build fipsEnabled

//go:debug fips140=only

package main

import (
	"context"
	"crypto/fips140"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/test/v4/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	log.Printf("FIPS enabled: %t", fips140.Enabled())
	if err := serve.Plugin(plugin.Plugin()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
