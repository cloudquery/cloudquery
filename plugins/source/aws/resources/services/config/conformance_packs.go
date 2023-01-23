package config

import (
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConformancePacks() *schema.Table {
	return &schema.Table{
		Name:        "aws_config_conformance_packs",
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_ConformancePackDetail.html`,
		Resolver:    fetchConfigConformancePacks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("config"),
		Transform:   transformers.TransformWithStruct(&types.ConformancePackDetail{}),
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
				Resolver: schema.PathResolver("ConformancePackArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			ConformancePackRuleCompliances(),
		},
	}
}
