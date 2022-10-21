package recipes

import (
	types "github.com/aws/aws-sdk-go-v2/service/identitystore/types"
)

func IdentitystoreResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "groups",
			Struct:     &types.Group{},
		},
		{
			SubService: "users",
			Struct:     &types.User{},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "identitystore"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("identitystore")`
	}
	return resources
}
