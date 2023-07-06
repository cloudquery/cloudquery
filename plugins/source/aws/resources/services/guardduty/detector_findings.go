package guardduty

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func detectorFindings() *schema.Table {
	tableName := "aws_guardduty_detector_findings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Finding.html`,
		Resolver:    fetchDetectorFindings,
		Transform: transformers.TransformWithStruct(&types.Finding{},
			transformers.WithTypeTransformer(client.TimestampTypeTransformer),
			transformers.WithResolverTransformer(client.TimestampResolverTransformer),
			transformers.WithPrimaryKeys("Arn"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "guardduty"),
		Columns: []schema.Column{
			{
				Name:       "detector_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchDetectorFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(*models.DetectorWrapper)

	cl := meta.(*client.Client)
	svc := cl.Services().Guardduty
	config := &guardduty.ListFindingsInput{
		DetectorId: &detector.Id,
	}
	paginator := guardduty.NewListFindingsPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *guardduty.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.FindingIds) == 0 {
			continue
		}

		f, err := svc.GetFindings(ctx, &guardduty.GetFindingsInput{
			DetectorId: &detector.Id,
			FindingIds: page.FindingIds,
		}, func(options *guardduty.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- f.Findings
	}
	return nil
}
