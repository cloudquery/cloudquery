//go:build linux && boringcrypto

package main

import (
	"context"
	_ "crypto/tls/fipsonly"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/test/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	p := serve.Plugin(plugin.Plugin())
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
