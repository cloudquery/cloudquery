package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchLightsailInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	input := lightsail.GetInstancesInput{}
	for {
		output, err := svc.GetInstances(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.Instances

		if aws.ToString(output.NextPageToken) == "" {
			break
		}
		input.PageToken = output.NextPageToken
	}
	return nil
}
func resolveLightsailInstanceAccessDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Instance)
	cli := meta.(*client.Client)
	svc := cli.Services().Lightsail
	input := lightsail.GetInstanceAccessDetailsInput{InstanceName: r.Name}
	output, err := svc.GetInstanceAccessDetails(ctx, &input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.AccessDetails)
}
func fetchLightsailInstancePortStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Instance)
	cli := meta.(*client.Client)
	svc := cli.Services().Lightsail
	input := lightsail.GetInstancePortStatesInput{InstanceName: r.Name}
	output, err := svc.GetInstancePortStates(ctx, &input)
	if err != nil {
		return err
	}

	res <- output.PortStates
	return nil
}
