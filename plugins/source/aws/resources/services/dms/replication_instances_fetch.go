package dms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type ReplicationInstanceWrapper struct {
	types.ReplicationInstance
	Tags map[string]interface{}
}

func fetchDmsReplicationInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().DMS

	var describeReplicationInstancesInput *databasemigrationservice.DescribeReplicationInstancesInput
	describeReplicationInstancesOutput, err := svc.DescribeReplicationInstances(ctx, describeReplicationInstancesInput)
	if err != nil {
		return err
	}
	if len(describeReplicationInstancesOutput.ReplicationInstances) == 0 {
		return nil
	}

	listTagsForResourceInput := databasemigrationservice.ListTagsForResourceInput{}
	for _, replicationInstance := range describeReplicationInstancesOutput.ReplicationInstances {
		listTagsForResourceInput.ResourceArnList = append(listTagsForResourceInput.ResourceArnList, *replicationInstance.ReplicationInstanceArn)
	}
	var listTagsForResourceOutput *databasemigrationservice.ListTagsForResourceOutput
	listTagsForResourceOutput, err = svc.ListTagsForResource(ctx, &listTagsForResourceInput)
	if err != nil {
		return err
	}
	replicationInstanceTags := make(map[string]map[string]interface{})
	for _, tag := range listTagsForResourceOutput.TagList {
		if replicationInstanceTags[*tag.ResourceArn] == nil {
			replicationInstanceTags[*tag.ResourceArn] = make(map[string]interface{})
		}
		replicationInstanceTags[*tag.ResourceArn][*tag.Key] = *tag.Value
	}

	for _, replicationInstance := range describeReplicationInstancesOutput.ReplicationInstances {
		wrapper := ReplicationInstanceWrapper{
			ReplicationInstance: replicationInstance,
			Tags:                replicationInstanceTags[*replicationInstance.ReplicationInstanceArn],
		}
		res <- wrapper
	}
	return nil
}
