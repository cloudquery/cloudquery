package dms

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/dms/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ReplicationInstances() *schema.Table {
	tableName := "aws_dms_replication_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html`,
		Resolver:    fetchDmsReplicationInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "dms"),
		Transform:   transformers.TransformWithStruct(&models.ReplicationInstanceWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ReplicationInstanceArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchDmsReplicationInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Databasemigrationservice

	var describeReplicationInstancesInput *databasemigrationservice.DescribeReplicationInstancesInput
	describeReplicationInstancesOutput, err := svc.DescribeReplicationInstances(ctx, describeReplicationInstancesInput, func(options *databasemigrationservice.Options) {
		options.Region = cl.Region
	})
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
	listTagsForResourceOutput, err = svc.ListTagsForResource(ctx, &listTagsForResourceInput, func(options *databasemigrationservice.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	replicationInstanceTags := make(map[string]map[string]any)
	for _, tag := range listTagsForResourceOutput.TagList {
		if replicationInstanceTags[*tag.ResourceArn] == nil {
			replicationInstanceTags[*tag.ResourceArn] = make(map[string]any)
		}
		replicationInstanceTags[*tag.ResourceArn][*tag.Key] = *tag.Value
	}

	for _, replicationInstance := range describeReplicationInstancesOutput.ReplicationInstances {
		wrapper := models.ReplicationInstanceWrapper{
			ReplicationInstance: replicationInstance,
			Tags:                replicationInstanceTags[*replicationInstance.ReplicationInstanceArn],
		}
		res <- wrapper
	}
	return nil
}
