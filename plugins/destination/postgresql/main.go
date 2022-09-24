package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/cloudquery/plugins/destinations/postgresql/client"
)

func main() {
	serve.Destination(client.New())
}

// https://github.com/cloudquery/cloudquery/releases/download/plugins-source-test-v1.1.0/test_darwin_arm64.zip
// https://github.com/cloudquery/cloudquery/releases/download/plugins-source-test-v1.1.5/test_darwin_arm64.zip