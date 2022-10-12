// Code generated by codegen; DO NOT EDIT.

package accessanalyzer

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AnalyzerArchiveRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_accessanalyzer_analyzer_archive_rules",
		Description: "https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_ArchiveRuleSummary.html",
		Resolver:    fetchAccessanalyzerAnalyzerArchiveRules,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "analyzer_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "filter",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Filter"),
			},
			{
				Name:     "rule_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RuleName"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}
