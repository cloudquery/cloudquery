package ec2

import (
	"context"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

var stateTransitionReasonTimeRegex = regexp.MustCompile(`\((.*)\)`)

func fetchEc2Instances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	p := ec2.NewDescribeInstancesPaginator(meta.(*client.Client).Services().Ec2,
		&ec2.DescribeInstancesInput{MaxResults: aws.Int32(1000)})

	for p.HasMorePages() {
		output, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, reservation := range output.Reservations {
			res <- reservation.Instances
		}
	}

	return nil
}

func resolveInstanceArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Instance)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "instance/" + aws.ToString(item.InstanceId),
	}
	return resource.Set(c.Name, a.String())
}

func resolveEc2InstanceStateTransitionReasonTime(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Item.(types.Instance)
	if instance.StateTransitionReason == nil {
		return nil
	}
	match := stateTransitionReasonTimeRegex.FindStringSubmatch(*instance.StateTransitionReason)
	if len(match) < 2 {
		// failed to get time from message
		return nil
	}
	const layout = "2006-01-02 15:04:05 MST"
	tm, err := time.Parse(layout, match[1])
	if err != nil {
		return err
	}
	return resource.Set(c.Name, tm)
}
