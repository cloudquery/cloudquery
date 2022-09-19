package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func KMSResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "keys",
			Struct:     &types.KeyMetadata{},
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rotation_enabled",
						Type:     schema.TypeBool,
						Resolver: `resolveKeysRotationEnabled`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveKeysTags`,
					},
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:          "replica_keys",
						Type:          schema.TypeJSON,
						Resolver:      `resolveKeysReplicaKeys`,
						IgnoreInTests: true,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "kms"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("kms")`
	}
	return resources
}
