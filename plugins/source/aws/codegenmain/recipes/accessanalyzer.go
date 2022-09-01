package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

var AccessAnalyzerResources = parentize(&Resource{
	DefaultColumns:       []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
	AWSStruct:            &types.AnalyzerSummary{},
	AWSService:           "AccessAnalyzer",
	CQSubserviceOverride: "accessanalyzers",
	Template:             "resource_get",
	ItemsStruct:          &accessanalyzer.ListAnalyzersOutput{},
	//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
},
	&Resource{
		AWSStruct:   &types.FindingSummary{},
		ItemsStruct: &accessanalyzer.ListFindingsOutput{},
		//ItemName:        "Finding",
		Template:        "resource_get",
		ParentFieldName: "Arn",
		ChildFieldName:  "AnalyzerArn",
	},
	&Resource{
		AWSStruct:   &types.ArchiveRuleSummary{},
		ItemsStruct: &accessanalyzer.ListArchiveRulesOutput{},
		//ItemName:        "ArchiveRules",
		Template:        "resource_get",
		ParentFieldName: "Name",
		ChildFieldName:  "AnalyzerName",
	},
)
