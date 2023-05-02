package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func detectorThreatIntelSets() *schema.Table {
	tableName := "aws_guardduty_detector_intel_sets"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetThreatIntelSet.html`,
		Resolver:            fetchDetectorThreatIntelSets,
		PreResourceResolver: getDetectorThreatIntelSet,
		Transform:           transformers.TransformWithStruct(&guardduty.GetThreatIntelSetOutput{}, transformers.WithPrimaryKeys("Name"), transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "guardduty"),
		Columns: []schema.Column{
			{
				Name:     "detector_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchDetectorThreatIntelSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(*models.DetectorWrapper)
	c := meta.(*client.Client)
	svc := c.Services().Guardduty
	config := &guardduty.ListThreatIntelSetsInput{DetectorId: aws.String(detector.Id)}
	paginator := guardduty.NewListThreatIntelSetsPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *guardduty.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- page.ThreatIntelSetIds
	}
	return nil
}

func getDetectorThreatIntelSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Guardduty
	id := resource.Item.(string)
	detector := resource.Parent.Item.(*models.DetectorWrapper)

	out, err := svc.GetThreatIntelSet(ctx, &guardduty.GetThreatIntelSetInput{
		DetectorId:       &detector.Id,
		ThreatIntelSetId: &id,
	}, func(options *guardduty.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}

	resource.Item = out
	return nil
}
