package frauddetector

import (
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Labels() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_labels",
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_Label.html`,
		Resolver:    fetchFrauddetectorLabels,
		Multiplex:   client.ServiceAccountRegionMultiplexer("frauddetector"),
		Transform:   transformers.TransformWithStruct(&types.Label{}),
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
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveResourceTags,
			},
		},
	}
}
