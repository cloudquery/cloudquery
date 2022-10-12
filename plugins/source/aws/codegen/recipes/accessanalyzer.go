package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccessAnalyzerResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "analyzers",
			Struct:      &types.AnalyzerSummary{},
			Description: "https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_AnalyzerSummary.html",
			SkipFields:  []string{"Arn"},
			Multiplex:   `client.ServiceAccountRegionMultiplexer("access-analyzer")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{"AnalyzerFindings()", "AnalyzerArchiveRules()"},
		},
		{
			SubService:  "analyzer_findings",
			Struct:      &types.FindingSummary{},
			Description: "https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_FindingSummary.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveFindingArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "analyzer_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "analyzer_archive_rules",
			Struct:      &types.ArchiveRuleSummary{},
			Description: "https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_ArchiveRuleSummary.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "analyzer_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "accessanalyzer"
	}
	return resources
}
