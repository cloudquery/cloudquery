package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FirehoseResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "delivery_streams",
			Struct:              &types.DeliveryStreamDescription{},
			Description:         "https://docs.aws.amazon.com/firehose/latest/APIReference/API_DeliveryStreamDescription.html",
			SkipFields:          []string{"DeliveryStreamARN"},
			PreResourceResolver: "getDeliveryStream",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveFirehoseDeliveryStreamTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DeliveryStreamARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "firehose"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("firehose")`
	}
	return resources
}
