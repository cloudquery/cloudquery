package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleanalytics/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v2/serve"
)

const sentryDSN = "https://1389ab4b050f4862879db19a783ecb78@o1396617.ingest.sentry.io/4504792343379968"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
