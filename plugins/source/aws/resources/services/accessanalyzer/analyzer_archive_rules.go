package accessanalyzer

import (
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AnalyzerArchiveRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_accessanalyzer_analyzer_archive_rules",
		Description: `https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_ArchiveRuleSummary.html`,
		Resolver:    fetchAccessanalyzerAnalyzerArchiveRules,
		Transform:   transformers.TransformWithStruct(&types.ArchiveRuleSummary{}),
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
		},
	}
}
