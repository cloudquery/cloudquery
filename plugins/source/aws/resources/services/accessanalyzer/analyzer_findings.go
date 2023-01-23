package accessanalyzer

import (
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AnalyzerFindings() *schema.Table {
	return &schema.Table{
		Name:        "aws_accessanalyzer_analyzer_findings",
		Description: `https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_FindingSummary.html`,
		Resolver:    fetchAccessanalyzerAnalyzerFindings,
		Transform:   transformers.TransformWithStruct(&types.FindingSummary{}),
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
				Resolver: resolveFindingArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "analyzer_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
