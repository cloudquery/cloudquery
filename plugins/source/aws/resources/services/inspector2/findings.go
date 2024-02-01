package inspector2

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Findings() *schema.Table {
	tableName := "aws_inspector2_findings"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/inspector/v2/APIReference/API_Finding.html

The ` + "`request_account_id` and `request_region` columns are added to show from where the request was made.",
		Resolver:  fetchFindings,
		Transform: transformers.TransformWithStruct(&types.Finding{}),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, client.AWSServiceInspector2.String()),
		Columns: schema.ColumnList{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("FindingArn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchFindings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceInspector2).Inspector2

	input := inspector2.ListFindingsInput{MaxResults: aws.Int32(100)}
	paginator := inspector2.NewListFindingsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *inspector2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Findings
	}

	return nil
}
