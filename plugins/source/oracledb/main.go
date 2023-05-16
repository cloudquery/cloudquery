package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracledb/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v2/serve"
)

const sentryDSN = "https://3613e006170b45399a66a93c992225e4@o1396617.ingest.sentry.io/4504922534445056"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
