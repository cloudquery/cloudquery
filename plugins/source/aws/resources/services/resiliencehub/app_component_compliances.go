package resiliencehub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func appComponentCompliances() *schema.Table {
	tableName := "aws_resiliencehub_app_component_compliances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppComponentCompliance.html`,
		Resolver:    fetchAppComponentCompliances,
		Transform:   transformers.TransformWithStruct(&types.AppComponentCompliance{}, transformers.WithPrimaryKeys("AppComponentName")),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, assessmentARN},
	}
}

func fetchAppComponentCompliances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceResiliencehub).Resiliencehub
	p := resiliencehub.NewListAppComponentCompliancesPaginator(svc, &resiliencehub.ListAppComponentCompliancesInput{AssessmentArn: parent.Item.(*types.AppAssessment).AppArn})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *resiliencehub.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- out.ComponentCompliances
	}
	return nil
}
