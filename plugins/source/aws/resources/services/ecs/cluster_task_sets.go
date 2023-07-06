package ecs

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func clusterTaskSets() *schema.Table {
	tableName := "aws_ecs_cluster_task_sets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskSet.html`,
		Resolver:    fetchEcsClusterTaskSets,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ecs"),
		Transform:   transformers.TransformWithStruct(&types.TaskSet{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("TaskSetArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
func fetchEcsClusterTaskSets(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	cluster := resource.Parent.Item.(types.Cluster)
	service := resource.Item.(types.Service)

	cl := meta.(*client.Client)
	svc := cl.Services().Ecs
	config := ecs.DescribeTaskSetsInput{
		Cluster: cluster.ClusterArn,
		Service: service.ServiceArn,
		Include: []types.TaskSetField{types.TaskSetFieldTags},
	}

	taskSets, err := svc.DescribeTaskSets(ctx, &config, func(options *ecs.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- taskSets.TaskSets
	return nil
}
