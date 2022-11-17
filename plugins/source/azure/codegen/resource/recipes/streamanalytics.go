package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/streamanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func StreamAnalytics() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armstreamanalytics.StreamingJob),
			Resolver: &resource.FuncParams{
				Func: streamanalytics.StreamingJobsClient.NewListPager,
			},
		},
	}
}
