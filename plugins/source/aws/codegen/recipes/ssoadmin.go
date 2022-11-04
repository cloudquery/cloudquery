package recipes

import (
	types "github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
)

func SSOAdminResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "instances",
			Struct:     &types.InstanceMetadata{},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "ssoadmin"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("identitystore")`
	}
	return resources
}
