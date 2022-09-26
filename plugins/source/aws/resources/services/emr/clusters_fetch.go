package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
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
		for _, c := range response.Clusters {
			out, err := svc.DescribeCluster(ctx, &emr.DescribeClusterInput{ClusterId: c.Id})
			if err != nil {
				return err
			}
			res <- out.Cluster
		}
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
