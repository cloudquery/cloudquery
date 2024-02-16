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

func detectorThreatIntelSets() *schema.Table {
	tableName := "aws_guardduty_detector_intel_sets"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetThreatIntelSet.html`,
		Resolver:            fetchDetectorThreatIntelSets,
		PreResourceResolver: getDetectorThreatIntelSet,
		Transform: transformers.TransformWithStruct(&models.ThreatIntelSetWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
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
						"threatintelset",
						resource.Item.(models.ThreatIntelSetWrapper).Id,
					}, nil
				}),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchDetectorThreatIntelSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(models.DetectorWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGuardduty).Guardduty
	config := &guardduty.ListThreatIntelSetsInput{DetectorId: aws.String(detector.Id)}
	paginator := guardduty.NewListThreatIntelSetsPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *guardduty.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ThreatIntelSetIds
	}
	return nil
}

func getDetectorThreatIntelSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGuardduty).Guardduty
	id := resource.Item.(string)
	detector := resource.Parent.Item.(models.DetectorWrapper)

	out, err := svc.GetThreatIntelSet(ctx, &guardduty.GetThreatIntelSetInput{
		DetectorId:       &detector.Id,
		ThreatIntelSetId: &id,
	}, func(options *guardduty.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = models.ThreatIntelSetWrapper{GetThreatIntelSetOutput: out, Id: id}
	return nil
}
