package elasticsearch

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func VpcEndpoints() *schema.Table {
	tableName := "aws_elasticsearch_vpc_endpoints"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_VpcEndpoint.html`,
		Resolver:    fetchElasticsearchVpcEndpoints,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "es"),
		Transform:   transformers.TransformWithStruct(&types.VpcEndpoint{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("VpcEndpointId"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchElasticsearchVpcEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticsearchservice
	// get the IDs first
	listInput := new(elasticsearchservice.ListVpcEndpointsInput)
	var vpcEndpointIDs []string
	// No paginator available
	for {
		out, err := svc.ListVpcEndpoints(ctx, listInput, func(options *elasticsearchservice.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		for _, summary := range out.VpcEndpointSummaryList {
			vpcEndpointIDs = append(vpcEndpointIDs, *summary.VpcEndpointId)
		}

		if aws.ToString(out.NextToken) == "" {
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
			func(options *elasticsearchservice.Options) {
				options.Region = cl.Region
			},
		)
		if err != nil {
			return err
		}

		res <- out.VpcEndpoints
	}

	return nil
}
