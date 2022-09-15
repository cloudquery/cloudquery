package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudfrontResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "cache_policies",
			Struct:     &types.CachePolicySummary{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{

					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveCachePolicyARN()`,
					},
				}...),
		},
		{
			SubService: "distributions",
			Struct:     &types.Distribution{},
			SkipFields: []string{"ARN"},
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
