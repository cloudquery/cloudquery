package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func XRayResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "encryption_configs",
			Struct:      &types.EncryptionConfig{},
			Description: "https://docs.aws.amazon.com/xray/latest/api/API_EncryptionConfig.html",
		},
		{
			SubService:  "groups",
			Struct:      &types.Group{},
			Description: "https://docs.aws.amazon.com/xray/latest/api/API_Group.html",
			SkipFields:  []string{"GroupARN"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("GroupARN")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `resolveXrayGroupTags`,
				},
			},
		},
		{
			SubService:  "sampling_rules",
			Struct:      &types.SamplingRuleRecord{},
			Description: "https://docs.aws.amazon.com/xray/latest/api/API_SamplingRuleRecord.html",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("SamplingRule.RuleARN")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `resolveXraySamplingRuleTags`,
				},
			},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "xray"
		r.ExtraColumns = append(defaultRegionalColumns, r.ExtraColumns...)
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("xray")`
	}
	return resources
}
