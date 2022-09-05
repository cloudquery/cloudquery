package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func init() {
	add(&Resource{
		DefaultColumns:         []codegen.ColumnDefinition{AccountIdColumn},
		AWSStruct:              &types.CachePolicySummary{},
		AWSService:             "Cloudfront",
		RawMultiplexerOverride: "client.AccountMultiplex",
		Template:               "resource_get",
		ItemsStruct:            &cloudfront.ListCachePoliciesOutput{},
		PageTokenInputField:    "Marker",
		ColumnOverrides: map[string]codegen.ColumnDefinition{
			"arn": {
				Type:     schema.TypeString,
				Resolver: "resolvers.ResolveCachePolicyArn",
			},
		},
	},
		&Resource{
			DefaultColumns:      []codegen.ColumnDefinition{AccountIdColumn},
			AWSStruct:           &types.Distribution{},
			AWSService:          "Cloudfront",
			Template:            "resource_list_describe",
			PaginatorStruct:     &cloudfront.ListDistributionsOutput{},
			PaginatorGetStruct:  &cloudfront.GetDistributionInput{},
			ItemsStruct:         &cloudfront.GetDistributionOutput{},
			PageTokenInputField: "Marker",
		},
	)
}
