package ecs

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func clusterTasks() *schema.Table {
	tableName := "aws_ecs_cluster_tasks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Task.html`,
		Resolver:    fetchEcsClusterTasks,
		Transform:   transformers.TransformWithStruct(&types.Task{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("TaskArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "task_protection",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: getEcsTaskProtection,
			},
		},
	}
}

func fetchEcsClusterTasks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(types.Cluster)

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEcs).Ecs
	input := &ecs.ListTasksInput{
		Cluster:    cluster.ClusterArn,
		MaxResults: aws.Int32(100),
	}

	paginator := ecs.NewListTasksPaginator(svc, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ecs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.TaskArns) == 0 {
			continue
		}
		describeServicesInput := ecs.DescribeTasksInput{
			Cluster: cluster.ClusterArn,
			Tasks:   page.TaskArns,
			Include: []types.TaskField{types.TaskFieldTags},
		}
		describeTasks, err := svc.DescribeTasks(ctx, &describeServicesInput, func(options *ecs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- describeTasks.Tasks
	}

	return nil
}

func getEcsTaskProtection(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEcs).Ecs
	task := resource.Item.(types.Task)
	resp, err := svc.GetTaskProtection(ctx, &ecs.GetTaskProtectionInput{
		Cluster: task.ClusterArn,
		Tasks:   []string{aws.ToString(task.TaskArn)},
	}, func(options *ecs.Options) {
		options.Region = cl.Region
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
