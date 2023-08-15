package guardduty

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Detectors() *schema.Table {
	tableName := "aws_guardduty_detectors"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetDetector.html`,
		Resolver:            fetchGuarddutyDetectors,
		PreResourceResolver: getDetector,
		Transform: transformers.TransformWithStruct(&models.DetectorWrapper{},
			transformers.WithTypeTransformer(client.TimestampTypeTransformer),
			transformers.WithResolverTransformer(client.TimestampResolverTransformer),
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "guardduty"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: resolveGuarddutyARN(),
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			detectorFindings(),
			detectorFilters(),
			detectorMembers(),
			detectorIPSets(),
			detectorPublishingDestinations(),
			detectorThreatIntelSets(),
		},
	}
}

func fetchGuarddutyDetectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Guardduty
	config := &guardduty.ListDetectorsInput{}
	paginator := guardduty.NewListDetectorsPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *guardduty.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DetectorIds
	}
	return nil
}

func getDetector(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Guardduty
	dId := resource.Item.(string)

	d, err := svc.GetDetector(ctx, &guardduty.GetDetectorInput{DetectorId: &dId}, func(options *guardduty.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = &models.DetectorWrapper{GetDetectorOutput: d, Id: dId}
	return nil
}

func resolveGuarddutyARN() schema.ColumnResolver {
	return client.ResolveARN(client.GuardDutyService, func(resource *schema.Resource) ([]string, error) {
		return []string{"detector", resource.Item.(*models.DetectorWrapper).Id}, nil
	})
}
