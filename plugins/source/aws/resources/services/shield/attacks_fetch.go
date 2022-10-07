package shield

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchShieldAttacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	end := time.Now()
	start := end.Add(-time.Hour * 24)
	config := shield.ListAttacksInput{
		EndTime:   &types.TimeRange{ToExclusive: &end},
		StartTime: &types.TimeRange{FromInclusive: &start},
	}
	for {
		output, err := svc.ListAttacks(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.AttackSummaries

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func getAttack(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	a := resource.Item.(types.AttackSummary)

	attack, err := svc.DescribeAttack(ctx, &shield.DescribeAttackInput{AttackId: a.AttackId})
	if err != nil {
		return err
	}

	resource.Item = attack.Attack
	return nil
}
