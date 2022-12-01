package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/dashboard_lists"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/dashboards"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/downtimes"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/hosts"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/incidents"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/monitors"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/notebooks"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/roles"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/synthetics"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/users"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Tables() []*schema.Table {
	return []*schema.Table{
		dashboard_lists.DashboardLists(),
		dashboards.Dashboards(),
		downtimes.Downtimes(),
		hosts.Hosts(),
		incidents.Incidents(),
		monitors.Monitors(),
		notebooks.Notebooks(),
		roles.Roles(),
		roles.Permissions(),
		synthetics.GlobalVariables(),
		synthetics.Synthetics(),
		users.Users(),
	}
}
