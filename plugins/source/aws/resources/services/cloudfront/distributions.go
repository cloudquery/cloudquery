package cloudfront

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Distributions() *schema.Table {
	return &schema.Table{
		Name:                "aws_cloudfront_distributions",
		Description:         `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_Distribution.html`,
		Resolver:            fetchCloudfrontDistributions,
		PreResourceResolver: getDistribution,
		Multiplex:           client.AccountMultiplex,
		Transform: 				 transformers.TransformWithStruct(&types.Distribution{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveCloudfrontDistributionTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
