package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEmrClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config emr.ListClustersInput
	c := meta.(*client.Client)
	svc := c.Services().EMR
	for {
		response, err := svc.ListClusters(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Clusters

		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func getCluster(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().EMR
	response, err := svc.DescribeCluster(ctx, &emr.DescribeClusterInput{ClusterId: resource.Item.(types.ClusterSummary).Id})
	if err != nil {
		return err
	}
	resource.Item = response.Cluster
	return nil
}
