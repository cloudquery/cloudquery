package guardduty

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func detectorIPSets() *schema.Table {
	tableName := "aws_guardduty_detector_ip_sets"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetIPSet.html`,
		Resolver:            fetchDetectorIPSets,
		PreResourceResolver: getDetectorIPSet,
		Transform: transformers.TransformWithStruct(&models.IPSetWrapper{},
			transformers.WithPrimaryKeyComponents("Name"),
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
						"ipset",
						resource.Item.(models.IPSetWrapper).Id,
					}, nil
				}),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchDetectorIPSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(models.DetectorWrapper)

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGuardduty).Guardduty
	config := &guardduty.ListIPSetsInput{
		DetectorId: &detector.Id,
	}
	paginator := guardduty.NewListIPSetsPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *guardduty.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.IpSetIds
	}
	return nil
}

func getDetectorIPSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGuardduty).Guardduty
	id := resource.Item.(string)
	detector := resource.Parent.Item.(models.DetectorWrapper)

	out, err := svc.GetIPSet(ctx, &guardduty.GetIPSetInput{
		DetectorId: &detector.Id,
		IpSetId:    &id,
	}, func(options *guardduty.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = models.IPSetWrapper{GetIPSetOutput: out, Id: id}
	return nil
}
