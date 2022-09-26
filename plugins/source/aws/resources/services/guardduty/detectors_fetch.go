package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type DetectorWrapper struct {
	*guardduty.GetDetectorOutput
	Id string
}

func fetchGuarddutyDetectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().GuardDuty
	config := &guardduty.ListDetectorsInput{}
	for {
		output, err := svc.ListDetectors(ctx, config)
		if err != nil {
			return err
		}
		for _, dId := range output.DetectorIds {
			d, err := svc.GetDetector(ctx, &guardduty.GetDetectorInput{DetectorId: aws.String(dId)}, func(o *guardduty.Options) {
				o.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- DetectorWrapper{d, dId}
		}
		if output.NextToken == nil {
			return nil
		}
		config.NextToken = output.NextToken
	}
}

func fetchGuarddutyDetectorMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	detector := parent.Item.(DetectorWrapper)
	c := meta.(*client.Client)
	svc := c.Services().GuardDuty
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
		return []string{"detector", resource.Item.(DetectorWrapper).Id}, nil
	})
}
