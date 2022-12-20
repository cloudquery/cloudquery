package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
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

	return describeVPCEndpoints(ctx, svc, res, vpcEndpointIDs)
}

func describeVPCEndpoints(ctx context.Context, svc services.ElasticsearchserviceClient, res chan<- any, vpcEndpointIDs []string) error {
	// DescribeVpcEndpoints supports amounts [5, 100], so, if we have > 100 endpoints to fetch, split in halves
	const maxLen = 100
	if l := len(vpcEndpointIDs); l > maxLen {
		if err := describeVPCEndpoints(ctx, svc, res, vpcEndpointIDs[:l/2]); err != nil {
			return err
		}
		return describeVPCEndpoints(ctx, svc, res, vpcEndpointIDs[l/2:])
	}

	out, err := svc.DescribeVpcEndpoints(ctx,
		&elasticsearchservice.DescribeVpcEndpointsInput{VpcEndpointIds: vpcEndpointIDs},
	)
	if err != nil {
		return err
	}

	res <- out.VpcEndpoints

	return nil
}
