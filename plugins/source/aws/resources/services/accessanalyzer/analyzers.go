package accessanalyzer

import (
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Analyzers() *schema.Table {
	return &schema.Table{
		Name:        "aws_accessanalyzer_analyzers",
		Description: `https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_AnalyzerSummary.html`,
		Resolver:    fetchAccessanalyzerAnalyzers,
		Multiplex:   client.ServiceAccountRegionMultiplexer("access-analyzer"),
		Transform:   transformers.TransformWithStruct(&types.AnalyzerSummary{}),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			AnalyzerFindings(),
			AnalyzerArchiveRules(),
		},
	}
}
