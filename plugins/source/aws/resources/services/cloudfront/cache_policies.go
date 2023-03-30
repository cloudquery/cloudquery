package cloudfront

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CachePolicies() *schema.Table {
	tableName := "aws_cloudfront_cache_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CachePolicySummary.html`,
		Resolver:    fetchCloudfrontCachePolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudfront"),
		Transform:   client.TransformWithStruct(&types.CachePolicySummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CachePolicy.Id"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveCachePolicyARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
