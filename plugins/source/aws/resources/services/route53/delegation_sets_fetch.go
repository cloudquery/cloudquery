package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRoute53DelegationSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config route53.ListReusableDelegationSetsInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	for {
		response, err := svc.ListReusableDelegationSets(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.DelegationSets
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func resolveDelegationSetArn() schema.ColumnResolver {
	return client.ResolveARNGlobal(client.Route53Service, func(resource *schema.Resource) ([]string, error) {
		return []string{"delegationset", *resource.Item.(types.DelegationSet).Id}, nil
	})
}
