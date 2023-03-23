package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Detectors() *schema.Table {
	tableName := "aws_guardduty_detectors"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetDetector.html`,
		Resolver:            fetchGuarddutyDetectors,
		PreResourceResolver: getDetector,
		Transform:           transformers.TransformWithStruct(&models.DetectorWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "guardduty"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGuarddutyARN(),
			},
			{
				Name: "id",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			detectorMembers(),
			detectorFindings(),
			detectorFilters(),
			detectorIPSets(),
		},
	}
}

func fetchGuarddutyDetectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Guardduty
	config := &guardduty.ListDetectorsInput{}
	for {
		output, err := svc.ListDetectors(ctx, config)
		if err != nil {
			return err
		}
		res <- output.DetectorIds

		if output.NextToken == nil {
			return nil
		}
		config.NextToken = output.NextToken
	}
}

func getDetector(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Guardduty
	dId := resource.Item.(string)

	d, err := svc.GetDetector(ctx, &guardduty.GetDetectorInput{DetectorId: &dId})
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
