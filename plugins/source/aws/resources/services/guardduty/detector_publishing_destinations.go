package guardduty

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func detectorPublishingDestinations() *schema.Table {
	tableName := "aws_guardduty_detector_publishing_destinations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DescribePublishingDestination.html`,
		Resolver:    fetchGuarddutyDetectorPublishingDestinations,
		Transform:   transformers.TransformWithStruct(&types.Destination{}, transformers.WithPrimaryKeys("DestinationId")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "guardduty"),
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

func fetchGuarddutyDetectorPublishingDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(*models.DetectorWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Guardduty
	config := &guardduty.ListPublishingDestinationsInput{DetectorId: aws.String(detector.Id)}
	paginator := guardduty.NewListPublishingDestinationsPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *guardduty.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Destinations
	}
	return nil
}
