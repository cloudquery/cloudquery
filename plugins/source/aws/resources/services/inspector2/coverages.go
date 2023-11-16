package inspector2

import (
	"context"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Coverages() *schema.Table {
	tableName := "aws_inspector2_coverages"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/inspector/v2/APIReference/API_ListCoverage.html
The 'request_account_id' and 'request_region' columns are added to show from where the request was made.`,
		Resolver:  fetchInspector2Coverages,
		Transform: transformers.TransformWithStruct(&types.CoveredResource{}, transformers.WithSkipFields("ResourceId")),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "inspector2"),
		Columns: []schema.Column{
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ResourceId"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchInspector2Coverages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceInspector2).Inspector2
	input := inspector2.ListCoverageInput{}
	paginator := inspector2.NewListCoveragePaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *inspector2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.CoveredResources
	}
	return nil
}
