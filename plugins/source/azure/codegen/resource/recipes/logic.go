package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/logic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/monitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Logic() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armlogic.Workflow),
			Resolver: &resource.FuncParams{
				Func: logic.WorkflowsClient.NewListBySubscriptionPager,
			},
			Children: []*resource.Resource{
				{
					SubService: "diagnostic_settings",
					Struct:     new(armmonitor.DiagnosticSettingsResource),
					Resolver: &resource.FuncParams{
						Func:   monitor.DiagnosticSettingsClient.NewListPager,
						Params: []string{"id.String()"},
					},
				},
			},
		},
	}
}
