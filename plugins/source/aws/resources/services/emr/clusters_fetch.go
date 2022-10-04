package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEmrClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return client.ListAndDetailResolver(ctx, meta, res, listEmrClusters, emrClusterDetail)
}

func listEmrClusters(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	var config emr.ListClustersInput
	c := meta.(*client.Client)
	svc := c.Services().EMR
	for {
		output, err := svc.ListClusters(ctx, &config)
		if err != nil {
			return err
		}
		for _, item := range output.Clusters {
			detailChan <- item.Id
		}
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func emrClusterDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	c := meta.(*client.Client)
	svc := c.Services().EMR

	clusterId := listInfo.(*string)

	out, err := svc.DescribeCluster(ctx, &emr.DescribeClusterInput{ClusterId: clusterId})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- err
		return
	}

	resultsChan <- out.Cluster
}
