package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/resources/services/auditlog"
	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/resources/services/projects"
	"github.com/cloudquery/plugin-sdk/v2/caser"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

var (
	Version = "development"
)

var customExceptions = map[string]string{
	"launchdarkly": "LaunchDarkly",
	"auditlog":     "Audit Log",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	exceptions := make(map[string]string)
	for k, v := range source.DefaultTitleExceptions {
		exceptions[k] = v
	}
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	return csr.ToTitle(table.Name)
}

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"launchdarkly",
		Version,
		[]*schema.Table{
			auditlog.AuditLogEntries(),
			projects.Projects(),
		},
		client.Configure,
		source.WithTitleTransformer(titleTransformer),
	)
}
