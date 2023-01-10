package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

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

func fetchGuarddutyDetectorMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(*models.DetectorWrapper)
	c := meta.(*client.Client)
	svc := c.Services().Guardduty
	config := &guardduty.ListMembersInput{DetectorId: aws.String(detector.Id)}
	for {
		output, err := svc.ListMembers(ctx, config)
		if err != nil {
			return err
		}
		res <- output.Members
		if output.NextToken == nil {
			return nil
		}
		config.NextToken = output.NextToken
	}
}

func resolveGuarddutyARN() schema.ColumnResolver {
	return client.ResolveARN(client.GuardDutyService, func(resource *schema.Resource) ([]string, error) {
		return []string{"detector", resource.Item.(*models.DetectorWrapper).Id}, nil
	})
}
