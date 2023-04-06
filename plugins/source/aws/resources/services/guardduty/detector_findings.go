package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func detectorFindings() *schema.Table {
	tableName := "aws_guardduty_detector_findings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Finding.html`,
		Resolver:    fetchDetectorFindings,
		Transform:   transformers.TransformWithStruct(&types.Finding{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "guardduty"),
		Columns: []schema.Column{
			{
				Name:     "detector_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchDetectorFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(*models.DetectorWrapper)

	c := meta.(*client.Client)
	svc := c.Services().Guardduty
	config := &guardduty.ListFindingsInput{
		DetectorId: &detector.Id,
	}
	for {
		output, err := svc.ListFindings(ctx, config)
		if err != nil {
			return err
		}
		if len(output.FindingIds) == 0 {
			return nil
		}

		f, err := svc.GetFindings(ctx, &guardduty.GetFindingsInput{
			DetectorId: &detector.Id,
			FindingIds: output.FindingIds,
		})
		if err != nil {
			return err
		}

		res <- f.Findings

		if aws.ToString(output.NextToken) == "" {
			return nil
		}
		config.NextToken = output.NextToken
	}
}
