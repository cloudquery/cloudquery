package ecs

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("TaskArn"),
				PrimaryKey: true,
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
	svc := cl.Services().Ecs
	var allConfigs []tableoptions.CustomListTasksOpts
	if cl.Spec.TableOptions.ECSTasks != nil && cl.Spec.TableOptions.ECSTasks.ListTasksOpts != nil {
		allConfigs = cl.Spec.TableOptions.ECSTasks.ListTasksOpts
	} else {
		allConfigs = []tableoptions.CustomListTasksOpts{{ListTasksInput: ecs.ListTasksInput{MaxResults: aws.Int32(100)}}}
	}
	for _, config := range allConfigs {
		config.Cluster = cluster.ClusterArn
		paginator := ecs.NewListTasksPaginator(svc, &config.ListTasksInput)
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
	}
	return nil
}

func getEcsTaskProtection(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecs
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
