package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TransferResources() []*Resource {
	resources := []*Resource{

		{
			SubService:          "servers",
			Struct:              &types.DescribedServer{},
			Description:         "https://docs.aws.amazon.com/transfer/latest/userguide/API_DescribedServer.html",
			SkipFields:          []string{"Arn", "Tags"},
			PreResourceResolver: "getServer",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "Specifies the key-value pairs that you can use to search for and group servers that were assigned to the server that was described",
						Type:        schema.TypeJSON,
						Resolver:    `resolveServersTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "transfer"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("transfer")`
	}
	return resources
}
