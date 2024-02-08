package guardduty

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
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
		Resolver:    fetchDetectorPublishingDestinations,
		Transform:   transformers.TransformWithStruct(&types.Destination{}, transformers.WithPrimaryKeyComponents("DestinationId")),
		Columns: schema.ColumnList{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			detectorARNColumn,
			{
				Name: "arn",
				Type: arrow.BinaryTypes.String,
				Resolver: client.ResolveARN(client.GuardDutyService, func(resource *schema.Resource) ([]string, error) {
					return []string{
						"detector",
						resource.Parent.Item.(models.DetectorWrapper).Id,
						"publishingDestination",
						aws.ToString(resource.Item.(types.Destination).DestinationId),
					}, nil
				}),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchDetectorPublishingDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(models.DetectorWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGuardduty).Guardduty
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
