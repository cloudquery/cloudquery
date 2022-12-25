package recipes

import (
	types "github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func IdentitystoreResources() []*Resource {
	resources := []*Resource{
		{
			TableDefinition: codegen.TableDefinition{
				SubService: "groups",
				Struct:     &types.Group{},
				Relations:  []string{"GroupMemberships()"},
			},
		},
		{
			TableDefinition: codegen.TableDefinition{
				SubService: "users",
				Struct:     &types.User{},
			},
		},
		{
			TableDefinition: codegen.TableDefinition{
				SubService: "group_memberships",
				Struct:     &types.GroupMembership{},
			},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "identitystore"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("identitystore")`
	}
	return resources
}
