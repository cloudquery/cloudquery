package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Cases() *schema.Table {
	tableName := "aws_support_cases"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeCases.html`,
		Resolver:    fetchCases,
		Transform:   transformers.TransformWithStruct(&types.CaseDetails{}, transformers.WithPrimaryKeys("CaseId")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "support"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
		Relations: []*schema.Table{communications()},
	}
}

func fetchCases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Support
	input := support.DescribeCasesInput{MaxResults: aws.Int32(100), IncludeResolvedCases: true}

	paginator := support.NewDescribeCasesPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *support.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Cases
	}

	return nil
}
