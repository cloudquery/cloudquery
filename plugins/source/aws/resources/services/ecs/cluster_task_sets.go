package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskSetArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
func fetchEcsClusterTaskSets(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	cluster := resource.Parent.Item.(types.Cluster)
	service := resource.Item.(types.Service)

	svc := meta.(*client.Client).Services().Ecs
	config := ecs.DescribeTaskSetsInput{
		Cluster: cluster.ClusterArn,
		Service: service.ServiceArn,
		Include: []types.TaskSetField{types.TaskSetFieldTags},
	}

	taskSets, err := svc.DescribeTaskSets(ctx, &config)
	if err != nil {
		return err
	}

	res <- taskSets.TaskSets
	return nil
}
