package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func analyzerFindings() *schema.Table {
	return &schema.Table{
		Name:        "aws_accessanalyzer_analyzer_findings",
		Description: `https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_FindingSummary.html`,
		Resolver:    fetchAccessanalyzerAnalyzerFindings,
		Transform:   transformers.TransformWithStruct(&types.FindingSummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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

func fetchAccessanalyzerAnalyzerFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	analyzer := parent.Item.(types.AnalyzerSummary)
	c := meta.(*client.Client)
	svc := c.Services().Accessanalyzer
	config := accessanalyzer.ListFindingsInput{
		AnalyzerArn: analyzer.Arn,
	}
	paginator := accessanalyzer.NewListFindingsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Findings
	}
	return nil
}

func resolveFindingArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "accessanalyzer",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "finding_summary/" + aws.ToString(resource.Item.(types.FindingSummary).Id),
	}
	return resource.Set(c.Name, a.String())
}
