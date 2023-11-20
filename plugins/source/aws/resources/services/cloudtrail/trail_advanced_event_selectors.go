package cloudtrail

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudtrail/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func trailAdvancedEventSelectors() *schema.Table {
	tableName := "aws_cloudtrail_trail_advanced_event_selectors"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedEventSelector.html`,
		Resolver:    fetchCloudtrailTrailAdvancedEventSelectors,
		Transform:   transformers.TransformWithStruct(&types.AdvancedEventSelector{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				// we can't use trail_arn as PK as single trail can have multiple advanced event selectors
				Name:     "trail_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchCloudtrailTrailAdvancedEventSelectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.CloudTrailWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudtrail).Cloudtrail
	response, err := svc.GetEventSelectors(ctx, &cloudtrail.GetEventSelectorsInput{TrailName: r.TrailARN}, func(options *cloudtrail.Options) {
		options.Region = *r.HomeRegion
	})
	if err != nil {
		return err
	}
	res <- response.AdvancedEventSelectors
	return nil
}
