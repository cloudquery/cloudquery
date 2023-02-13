package s3

import (
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AccessPoints() *schema.Table {
	return &schema.Table{
		Name:      "aws_s3_access_points",
		Resolver:  fetchAccessPoints,
		Transform: transformers.TransformWithStruct(&types.AccessPoint{}),
		Multiplex: client.ServiceAccountRegionMultiplexer("s3-control"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessPointArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
