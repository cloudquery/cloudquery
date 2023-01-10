package inspector2

import (
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Findings() *schema.Table {
	return &schema.Table{
		Name:        "aws_inspector2_findings",
		Description: `https://docs.aws.amazon.com/inspector/v2/APIReference/API_Finding.html`,
		Resolver:    fetchInspector2Findings,
		Transform:   transformers.TransformWithStruct(&types.Finding{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("inspector2"),
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
				Resolver: schema.PathResolver("FindingArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
