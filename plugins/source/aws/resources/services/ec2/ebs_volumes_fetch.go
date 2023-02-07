package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEc2EbsVolumes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	config := ec2.DescribeVolumesInput{}
	for {
		response, err := svc.DescribeVolumes(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Volumes
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveVolumeARN(partition string, region string, accountID string, volumeId string) string {
	a := arn.ARN{
		Partition: partition,
		Service:   "ec2",
		Region:    region,
		AccountID: accountID,
		Resource:  "volume/" + volumeId,
	}
	return a.String()
}

func resolveEbsVolumeArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	volume := resource.Item.(types.Volume)
	a := resolveVolumeARN(cl.Partition, cl.Region, cl.AccountID, aws.ToString(volume.VolumeId))
	return resource.Set(c.Name, a)
}
