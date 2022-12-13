package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://81e8135dfd0e42629810a434e0dd72cb@o1396617.ingest.sentry.io/4504317272915968"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
