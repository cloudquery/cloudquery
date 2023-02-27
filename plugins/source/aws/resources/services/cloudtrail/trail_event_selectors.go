package cloudtrail

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TrailEventSelectors() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudtrail_trail_event_selectors",
		Description: `https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_EventSelector.html`,
		Resolver:    fetchCloudtrailTrailEventSelectors,
		Multiplex:   client.ServiceAccountRegionMultiplexer("cloudtrail"),
		Transform:   transformers.TransformWithStruct(&types.EventSelector{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "trail_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
