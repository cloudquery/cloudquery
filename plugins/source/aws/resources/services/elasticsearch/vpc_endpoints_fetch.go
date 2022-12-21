package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchElasticsearchVpcEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Elasticsearchservice

	listInput := new(elasticsearchservice.ListVpcEndpointsInput)
	var vpcEndpointIDs []string
	// get the IDs first
	for {
		out, err := svc.ListVpcEndpoints(ctx, listInput)
		if err != nil {
			return err
		}

		for _, summary := range out.VpcEndpointSummaryList {
			vpcEndpointIDs = append(vpcEndpointIDs, *summary.VpcEndpointId)
		}

		if out.NextToken == nil {
			break
		}

		listInput.NextToken = out.NextToken
	}

	// slice in parts
	const maxLen = 100
	for len(vpcEndpointIDs) > 0 {
		var part []string
		if len(vpcEndpointIDs) > maxLen {
			part, vpcEndpointIDs = vpcEndpointIDs[:maxLen], vpcEndpointIDs[maxLen:]
		} else {
			part, vpcEndpointIDs = vpcEndpointIDs, nil
		}

		out, err := svc.DescribeVpcEndpoints(ctx,
			&elasticsearchservice.DescribeVpcEndpointsInput{VpcEndpointIds: part},
		)
		if err != nil {
			return err
		}

		res <- out.VpcEndpoints
	}

	return nil
}
