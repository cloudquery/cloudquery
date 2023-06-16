package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Operations() *schema.Table {
	tableName := "aws_route53_operations"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html`,
		Resolver:            fetchRoute53Operations,
		PreResourceResolver: getOperation,
		Transform:           transformers.TransformWithStruct(&route53domains.GetOperationDetailOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithPrimaryKeys("OperationId", "Status", "SubmittedDate", "Type")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "route53domains"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchRoute53Operations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Route53domains
	var input route53domains.ListOperationsInput
	paginator := route53domains.NewListOperationsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *route53domains.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Operations
	}
	return nil
}
func getOperation(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Route53domains
	v := resource.Item.(types.OperationSummary)

	d, err := svc.GetOperationDetail(ctx, &route53domains.GetOperationDetailInput{OperationId: v.OperationId}, func(options *route53domains.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = d

	return nil
}
