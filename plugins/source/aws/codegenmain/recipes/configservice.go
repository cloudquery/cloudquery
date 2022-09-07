package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	resolvers "github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/configservice"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func init() {
	add(combine(parentize(&Resource{
		DefaultColumns:             []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:                  &types.ConfigurationRecorder{},
		AWSService:                 "ConfigService",
		TablePrefixOverride:        "_",
		Template:                   "resource_get",
		MultiplexerServiceOverride: "config",
		ItemsStruct:                &configservice.DescribeConfigurationRecordersOutput{},
		PrimaryKeys:                []string{"arn"},
		ColumnOverrides: map[string]codegen.ColumnDefinition{
			"arn": {
				Type:     schema.TypeString,
				Resolver: "resolvers.ResolveConfigRecorderArn",
			},
		},
	},
		&Resource{
			// FIXME this used to be wrapped inside ConfigurationRecorder and fetched in a batch fashion
			AWSStruct:       &types.ConfigurationRecorderStatus{},
			Template:        "resource_get",
			ItemsStruct:     &configservice.DescribeConfigurationRecorderStatusOutput{},
			ParentFieldName: "[]string{*$.Name}",
			ChildFieldName:  "ConfigurationRecorderNames",
		},
	),
		parentize(&Resource{
			DefaultColumns:             []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
			AWSStruct:                  &types.ConformancePackDetail{},
			AWSService:                 "ConfigService",
			CQSubserviceOverride:       "conformance_packs",
			TablePrefixOverride:        "_",
			Template:                   "resource_get",
			MultiplexerServiceOverride: "config",
			ItemsStruct:                &configservice.DescribeConformancePacksOutput{},
			PrimaryKeys:                []string{"arn"},
			ColumnOverrides: map[string]codegen.ColumnDefinition{
				"conformance_pack_arn": {
					Name: "arn",
				},
			},
			CustomErrorBlock: `
		// This is a workaround until this bug is fixed = https://github.com/aws/aws-sdk-go-v2/issues/1539
		var ae smithy.APIError
		if (cl.Region == "af-south-1" || cl.Region == "ap-northeast-3") && errors.As(err, &ae) && ae.ErrorCode() == "AccessDeniedException" {
			return nil
		}`,
			Imports: []string{
				"github.com/aws/smithy-go",
			},
		},

			&Resource{
				AWSStruct: &resolvers.ConformancePackComplianceWrapper{},
				//TablePrefixOverride: "_",
				AWSSubService: "rule_compliances",
				Template:      "resource_manual",
				RawResolver:   "resolvers.FetchConformancePackRuleCompliances",
			},
		),
	)...)
}
