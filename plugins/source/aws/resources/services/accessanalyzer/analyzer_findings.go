package accessanalyzer

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func analyzerFindings() *schema.Table {
	return &schema.Table{
		Name:        "aws_accessanalyzer_analyzer_findings",
		Description: `https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_FindingSummary.html`,
		Resolver:    fetchAccessanalyzerAnalyzerFindings,
		Transform:   transformers.TransformWithStruct(&types.FindingSummary{}, transformers.WithPrimaryKeyComponents("Id")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "analyzer_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchAccessanalyzerAnalyzerFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	analyzer := parent.Item.(types.AnalyzerSummary)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAccessanalyzer).Accessanalyzer

	input := &accessanalyzer.ListFindingsInput{AnalyzerArn: analyzer.Arn}
	paginator := accessanalyzer.NewListFindingsPaginator(svc, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *accessanalyzer.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Findings
	}
	return nil
}
