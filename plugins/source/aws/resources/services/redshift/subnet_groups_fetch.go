package redshift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRedshiftSubnetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config redshift.DescribeClusterSubnetGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Redshift
	for {
		response, err := svc.DescribeClusterSubnetGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.ClusterSubnetGroups
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func resolveSubnetGroupArn() schema.ColumnResolver {
	return client.ResolveARN(client.RedshiftService, func(resource *schema.Resource) ([]string, error) {
		return []string{fmt.Sprintf("subnetgroup:%s", *resource.Item.(types.ClusterSubnetGroup).ClusterSubnetGroupName)}, nil
	})
}
