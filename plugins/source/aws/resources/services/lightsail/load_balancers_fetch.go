package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchLightsailLoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input lightsail.GetLoadBalancersInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetLoadBalancers(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.LoadBalancers
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func fetchLightsailLoadBalancerTlsCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.LoadBalancer)
	input := lightsail.GetLoadBalancerTlsCertificatesInput{
		LoadBalancerName: r.Name,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	response, err := svc.GetLoadBalancerTlsCertificates(ctx, &input)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- response.TlsCertificates
	return nil
}
