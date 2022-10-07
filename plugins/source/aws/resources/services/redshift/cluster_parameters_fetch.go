package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRedshiftClusterParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	group := parent.Item.(types.ClusterParameterGroupStatus)
	c := meta.(*client.Client)
	svc := c.Services().Redshift

	config := redshift.DescribeClusterParametersInput{
		ParameterGroupName: group.ParameterGroupName,
	}
	for {
		response, err := svc.DescribeClusterParameters(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Parameters
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}

	return nil
}
