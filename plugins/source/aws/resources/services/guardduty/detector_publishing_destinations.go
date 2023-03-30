package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func detectorPublishingDestinations() *schema.Table {
	tableName := "aws_guardduty_detector_publishing_destinations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Member.html`,
		Resolver:    fetchGuarddutyDetectorPublishingDestinations,
		Transform:   client.TransformWithStruct(&types.Destination{}, transformers.WithPrimaryKeys("DestinationId")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "guardduty"),
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

func fetchGuarddutyDetectorPublishingDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(*models.DetectorWrapper)
	c := meta.(*client.Client)
	svc := c.Services().Guardduty
	config := &guardduty.ListPublishingDestinationsInput{DetectorId: aws.String(detector.Id)}
	for {
		output, err := svc.ListPublishingDestinations(ctx, config)
		if err != nil {
			return err
		}
		res <- output.Destinations
		if output.NextToken == nil {
			return nil
		}
		config.NextToken = output.NextToken
	}
}
