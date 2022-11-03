package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudfrontResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "cache_policies",
			Struct:      &types.CachePolicySummary{},
			Description: "https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CachePolicySummary.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("CachePolicy.Id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveCachePolicyARN()`,
					},
				}...),
		},
		{
			SubService:          "distributions",
			Struct:              &types.Distribution{},
			Description:         "https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_Distribution.html",
			SkipFields:          []string{"ARN"},
			PreResourceResolver: "getDistribution",
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveCloudfrontDistributionTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "cloudfront"
		r.Multiplex = `client.AccountMultiplex`
	}
	return resources
}
