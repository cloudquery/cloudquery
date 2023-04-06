package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func clusterTasks() *schema.Table {
	tableName := "aws_ecs_cluster_tasks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Task.html`,
		Resolver:    fetchEcsClusterTasks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ecs"),
		Transform:   transformers.TransformWithStruct(&types.Task{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "task_protection",
				Type:     schema.TypeJSON,
				Resolver: getEcsTaskProtection,
			},
		},
	}
}

func fetchEcsClusterTasks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(types.Cluster)

	svc := meta.(*client.Client).Services().Ecs
	config := ecs.ListTasksInput{
		Cluster: cluster.ClusterArn,
	}
	for {
		listTasks, err := svc.ListTasks(ctx, &config)
		if err != nil {
			return err
		}
		if len(listTasks.TaskArns) == 0 {
			return nil
		}
		describeServicesInput := ecs.DescribeTasksInput{
			Cluster: cluster.ClusterArn,
			Tasks:   listTasks.TaskArns,
			Include: []types.TaskField{types.TaskFieldTags},
		}
		describeTasks, err := svc.DescribeTasks(ctx, &describeServicesInput)
		if err != nil {
			return err
		}

		res <- describeTasks.Tasks

		if aws.ToString(listTasks.NextToken) == "" {
			break
		}
		config.NextToken = listTasks.NextToken
	}
	return nil
}

func getEcsTaskProtection(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Ecs
	task := resource.Item.(types.Task)
	resp, err := svc.GetTaskProtection(ctx, &ecs.GetTaskProtectionInput{
		Cluster: task.ClusterArn,
		Tasks:   []string{aws.ToString(task.TaskArn)},
	})
	if err != nil {
		return err
	}
	if len(resp.Failures) > 0 {
		// This indicates that a task has been deleted in between the listing time and now
		return nil
	}
	return resource.Set(c.Name, resp.ProtectedTasks)
}
