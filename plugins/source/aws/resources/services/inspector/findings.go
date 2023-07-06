package inspector

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Findings() *schema.Table {
	tableName := "aws_inspector_findings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/inspector/v1/APIReference/API_Finding.html`,
		Resolver:    fetchInspectorFindings,
		Transform:   transformers.TransformWithStruct(&types.Finding{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "inspector"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
			{
				Name:     "attributes",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTagField("Attributes"),
			},
			{
				Name:     "user_attributes",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTagField("UserAttributes"),
			},
		},
	}
}

func fetchInspectorFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Inspector
	input := inspector.ListFindingsInput{MaxResults: aws.Int32(50)}
	paginator := inspector.NewListFindingsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *inspector.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.FindingArns) == 0 {
			continue
		}

		batch := 10
		for i := 0; i < len(page.FindingArns); i += batch {
			j := i + batch
			if j >= len(page.FindingArns) {
				j = len(page.FindingArns) - 1
			}
			out, err := svc.DescribeFindings(ctx, &inspector.DescribeFindingsInput{FindingArns: page.FindingArns[i:j]}, func(options *inspector.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return err
			}
			res <- out.Findings
		}
	}
	return nil
}
