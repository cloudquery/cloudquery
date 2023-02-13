package directconnect

import (
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:        "aws_directconnect_locations",
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_Location.html`,
		Resolver:    fetchDirectConnectLocations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("directconnect"),
		Transform:   transformers.TransformWithStruct(&types.Location{}, transformers.WithPrimaryKeys("LocationCode")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
