package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://29efd1efe53847c4a979bc2d2e712a74@o1396617.ingest.sentry.io/4504797528981504"

func main() {
	serve.Source(plugin.Plugin(),
		serve.WithSourceSentryDSN(sentryDSN),
	)
}
