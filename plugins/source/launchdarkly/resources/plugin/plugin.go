package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/resources/services/auditlog"
	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/resources/services/projects"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"launchdarkly",
		Version,
		[]*schema.Table{
			auditlog.AuditLogEntries(),
			projects.Projects(),
		},
		client.Configure,
	)
}
