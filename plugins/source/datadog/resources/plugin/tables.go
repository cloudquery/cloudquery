package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/dashboards"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/downtimes"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/hosts"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/incidents"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/monitors"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/notebooks"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/roles"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/rum"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/slos"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/synthetics"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/users"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"golang.org/x/exp/maps"
)

func titleTransformer(table *schema.Table) error {
	if table.Title != "" {
		return nil
	}
	exceptions := maps.Clone(docs.DefaultTitleExceptions)
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	table.Title = csr.ToTitle(table.Name)
	return nil
}

func getTables() schema.Tables {
	tables := []*schema.Table{
		dashboards.Dashboards(),
		dashboards.Lists(),
		downtimes.Downtimes(),
		hosts.Hosts(),
		incidents.Incidents(),
		monitors.Monitors(),
		notebooks.Notebooks(),
		roles.Roles(),
		roles.Permissions(),
		rum.Events(),
		slos.Objectives(),
		slos.Corrections(),
		synthetics.GlobalVariables(),
		synthetics.Synthetics(),
		users.Users(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, titleTransformer); err != nil {
		panic(err)
	}
	for _, table := range tables {
		schema.AddCqIDs(table)
	}
	return tables
}
