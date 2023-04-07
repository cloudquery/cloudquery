package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Analyzers() *schema.Table {
	tableName := "aws_accessanalyzer_analyzers"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_AnalyzerSummary.html`,
		Resolver:    fetchAccessanalyzerAnalyzers,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "access-analyzer"),
		Transform:   transformers.TransformWithStruct(&types.AnalyzerSummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
			analyzerFindings(),
			analyzerArchiveRules(),
		},
	}
}

func fetchAccessanalyzerAnalyzers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Accessanalyzer
	paginator := accessanalyzer.NewListAnalyzersPaginator(svc, &accessanalyzer.ListAnalyzersInput{})
	for paginator.HasMorePages() {
		// no need to override API call options anymore: https://github.com/aws/aws-sdk-go-v2/issues/1260
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Analyzers
	}
	return nil
}
