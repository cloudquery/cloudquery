package cloudfront

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Distributions() *schema.Table {
	tableName := "aws_cloudfront_distributions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_Distribution.html`,
		Resolver:            fetchCloudfrontDistributions,
		PreResourceResolver: getDistribution,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cloudfront"),
		Transform:           client.TransformWithStruct(&types.Distribution{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
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
