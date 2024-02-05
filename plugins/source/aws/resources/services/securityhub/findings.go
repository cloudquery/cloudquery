package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Findings() *schema.Table {
	tableName := "aws_securityhub_findings"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetFindings.html
` + "The `request_account_id` and `request_region` columns are added to show the account and region of where the request was made from." + `
This is useful when multi region and account aggregation is enabled.`,
		Resolver: fetchFindings,
		Transform: transformers.TransformWithStruct(&types.AwsSecurityFinding{},
			transformers.WithTypeTransformer(client.TimestampTypeTransformer),
			transformers.WithResolverTransformer(client.TimestampResolverTransformer),
			transformers.WithPrimaryKeyComponents("AwsAccountId", "Region", "CreatedAt", "UpdatedAt", "Description", "GeneratorId", "Id", "ProductArn", "SchemaVersion", "Title"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "securityhub"),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
		},
	}
}

func fetchFindings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSecurityhub).Securityhub

	config := securityhub.GetFindingsInput{MaxResults: aws.Int32(100)}
	p := securityhub.NewGetFindingsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(o *securityhub.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Findings
	}

	return nil
}
