package xray

import (
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:        "aws_xray_groups",
		Description: `https://docs.aws.amazon.com/xray/latest/api/API_Group.html`,
		Resolver:    fetchXrayGroups,
		Transform:   transformers.TransformWithStruct(&types.GroupSummary{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("xray"),
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
				Resolver: schema.PathResolver("GroupARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveXrayGroupTags,
			},
		},
	}
}
