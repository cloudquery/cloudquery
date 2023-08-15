package inspector2

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Findings() *schema.Table {
	tableName := "aws_inspector2_findings"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/inspector/v2/APIReference/API_Finding.html
The 'request_account_id' and 'request_region' columns are added to show from where the request was made.`,
		Resolver:  fetchInspector2Findings,
		Transform: transformers.TransformWithStruct(&types.Finding{}),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "inspector2"),
		Columns: []schema.Column{
			{
				Name:       "request_account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSAccount,
				PrimaryKey: true,
			},
			{
				Name:       "request_region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSRegion,
				PrimaryKey: true,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("FindingArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchInspector2Findings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Inspector2

	allConfigs := []tableoptions.CustomInspector2ListFindingsInput{{}}
	if cl.Spec.TableOptions.Inspector2Findings != nil {
		allConfigs = cl.Spec.TableOptions.Inspector2Findings.ListFindingsOpts
	}
	for _, input := range allConfigs {
		if input.MaxResults == nil {
			input.MaxResults = aws.Int32(100)
		}

		paginator := inspector2.NewListFindingsPaginator(svc, &input.ListFindingsInput)
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx, func(options *inspector2.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- page.Findings
		}
	}
	return nil
}
