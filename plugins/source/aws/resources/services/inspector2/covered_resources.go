package inspector2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func CoveredResources() *schema.Table {
	tableName := "aws_inspector2_covered_resources"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/inspector/v2/APIReference/API_CoveredResource.html

The ` + "`request_account_id` and `request_region` columns are added to show from where the request was made.",
		Resolver:  fetchCoveredResources,
		Transform: transformers.TransformWithStruct(&types.CoveredResource{}, transformers.WithPrimaryKeyComponents("AccountId", "ResourceId")),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, client.AWSServiceInspector2.String()),
		Columns: schema.ColumnList{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
		},
	}
}

func fetchCoveredResources(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceInspector2).Inspector2
	input := inspector2.ListCoverageInput{MaxResults: aws.Int32(200)}
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
