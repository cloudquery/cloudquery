package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_images",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Image.html`,
		Resolver:    fetchAppstreamImages,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appstream2"),
		Transform: transformers.TransformWithStruct(&types.Image{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
