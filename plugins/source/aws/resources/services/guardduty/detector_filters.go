package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func detectorFilters() *schema.Table {
	tableName := "aws_guardduty_detector_filters"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetFilter.html`,
		Resolver:            fetchDetectorFilters,
		PreResourceResolver: getDetectorFilter,
		Transform:           transformers.TransformWithStruct(&guardduty.GetFilterOutput{}, transformers.WithPrimaryKeys("Name"), transformers.WithSkipFields("ResultMetadata")),
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

func fetchDetectorFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(*models.DetectorWrapper)

	c := meta.(*client.Client)
	svc := c.Services().Guardduty
	config := &guardduty.ListFiltersInput{
		DetectorId: &detector.Id,
	}
	for {
		output, err := svc.ListFilters(ctx, config)
		if err != nil {
			return err
		}
		res <- output.FilterNames

		if output.NextToken == nil {
			return nil
		}
		config.NextToken = output.NextToken
	}
}

func getDetectorFilter(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Guardduty
	filterName := resource.Item.(string)
	detector := resource.Parent.Item.(*models.DetectorWrapper)

	out, err := svc.GetFilter(ctx, &guardduty.GetFilterInput{
		DetectorId: &detector.Id,
		FilterName: &filterName,
	})
	if err != nil {
		return err
	}

	resource.Item = out
	return nil
}
