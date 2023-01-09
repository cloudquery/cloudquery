package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Applications() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_applications",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Application.html`,
		Resolver:    fetchAppstreamApplications,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appstream2"),
		Transform:   transformers.TransformWithStruct(&types.Application{}),
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
		Relations: []*schema.Table{
			ApplicationFleetAssociations(),
		},
	}
}
