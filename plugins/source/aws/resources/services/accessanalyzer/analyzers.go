package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Analyzers() *schema.Table {
	tableName := "aws_accessanalyzer_analyzers"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_AnalyzerSummary.html`,
		Resolver:    fetchAccessanalyzerAnalyzers,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "access-analyzer"),
		Transform:   transformers.TransformWithStruct(&types.AnalyzerSummary{}, transformers.WithPrimaryKeyComponents("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
		Relations: []*schema.Table{
			analyzerFindings(),
			analyzerArchiveRules(),
		},
	}
}

func fetchAccessanalyzerAnalyzers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAccessanalyzer).Accessanalyzer
	paginator := accessanalyzer.NewListAnalyzersPaginator(svc, &accessanalyzer.ListAnalyzersInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *accessanalyzer.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Analyzers
	}
	return nil
}
