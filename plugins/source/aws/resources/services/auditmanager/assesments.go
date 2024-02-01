package auditmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Assessments() *schema.Table {
	tableName := "aws_auditmanager_assessments"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_Assessment.html`,
		Resolver:            fetchAssessments,
		PreResourceResolver: getAssessment,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "auditmanager"),
		Transform:           transformers.TransformWithStruct(&types.Assessment{}, transformers.WithPrimaryKeyComponents("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAuditmanager).Auditmanager
	paginator := auditmanager.NewListAssessmentsPaginator(svc, nil)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *auditmanager.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.AssessmentMetadata
	}
	return nil
}

func getAssessment(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAuditmanager).Auditmanager
	input := auditmanager.GetAssessmentInput{AssessmentId: resource.Item.(types.AssessmentMetadataItem).Id}

	output, err := svc.GetAssessment(ctx, &input, func(o *auditmanager.Options) { o.Region = cl.Region })
	if err != nil {
		return err
	}
	resource.Item = output.Assessment
	return nil
}
