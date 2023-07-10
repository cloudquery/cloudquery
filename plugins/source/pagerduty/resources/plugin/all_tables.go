package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/addons"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/business_services"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/escalation_policies"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/extension_schemas"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/extensions"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/incidents"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/maintenance_windows"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/priorities"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/rulesets"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/schedules"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/services"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/tags"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/teams"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/users"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/services/vendors"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func getTables() []*schema.Table {
	t := schema.Tables{
		addons.Addons(),
		incidents.Incidents(),
		business_services.BusinessServices(),
		escalation_policies.EscalationPolicies(),
		extension_schemas.ExtensionSchemas(),
		extensions.Extensions(),
		maintenance_windows.MaintenanceWindows(),
		priorities.Priorities(),
		rulesets.Rulesets(),
		schedules.Schedules(),
		services.Services(),
		tags.Tags(),
		teams.Teams(),
		users.Users(),
		vendors.Vendors(),
	}
	if err := transformers.TransformTables(t); err != nil {
		panic(err)
	}
	for _, table := range t {
		schema.AddCqIDs(table)
	}
	err := transformers.Apply(t, titleTransformer)
	if err != nil {
		panic(err)
	}
	return t
}
