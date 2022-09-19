package applicationautoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func applicationAutoscalingMultiplexer(multiplexer func(meta schema.ClientMeta) []schema.ClientMeta)  func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		clients := multiplexer(meta)
		var namespace types.ServiceNamespace
		namespaces := namespace.Values()
		allClients := make([]schema.ClientMeta, 0, len(clients)*len(namespaces))
		for _, c := range clients {
			for _, namespace := range namespaces {
				c := c.(*client.Client)
				c.AutoscalingNamespace = namespace
				allClients = append(allClients, c)
			}
		}
		return clients
	}
}

func fetchApplicationautoscalingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().ApplicationAutoscaling

	config := applicationautoscaling.DescribeScalingPoliciesInput{
		ServiceNamespace: types.ServiceNamespace(c.AutoscalingNamespace),
	}
	for {
		output, err := svc.DescribeScalingPolicies(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.ScalingPolicies

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
