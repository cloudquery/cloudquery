package guardduty

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func detectorFilters() *schema.Table {
	tableName := "aws_guardduty_detector_filters"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetFilter.html`,
		Resolver:            fetchDetectorFilters,
		PreResourceResolver: getDetectorFilter,
		Transform: transformers.TransformWithStruct(&guardduty.GetFilterOutput{},
			transformers.WithPrimaryKeyComponents("Name"),
			transformers.WithSkipFields("ResultMetadata"),
		),
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
						"filter",
						aws.ToString(resource.Item.(*guardduty.GetFilterOutput).Name),
					}, nil
				}),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchDetectorFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(models.DetectorWrapper)

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGuardduty).Guardduty
	config := &guardduty.ListFiltersInput{
		DetectorId: &detector.Id,
	}
	paginator := guardduty.NewListFiltersPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *guardduty.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.FilterNames
	}
	return nil
}

func getDetectorFilter(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGuardduty).Guardduty
	filterName := resource.Item.(string)
	detector := resource.Parent.Item.(models.DetectorWrapper)

	out, err := svc.GetFilter(ctx, &guardduty.GetFilterInput{
		DetectorId: &detector.Id,
		FilterName: &filterName,
	}, func(options *guardduty.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = out
	return nil
}
