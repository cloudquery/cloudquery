package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEcsClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ecs.ListClustersInput
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	for {
		listClustersOutput, err := svc.ListClusters(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		if len(listClustersOutput.ClusterArns) == 0 {
			return nil
		}
		describeClusterOutput, err := svc.DescribeClusters(ctx, &ecs.DescribeClustersInput{
			Clusters: listClustersOutput.ClusterArns,
			Include:  []types.ClusterField{types.ClusterFieldTags},
		}, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		res <- describeClusterOutput.Clusters

		if listClustersOutput.NextToken == nil {
			break
		}
		config.NextToken = listClustersOutput.NextToken
	}
	return nil
}

func resolveClustersSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster, ok := resource.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("expected to have types.Cluster but got %T", resource.Item)
	}
	settings := make(map[string]*string)
	for _, s := range cluster.Settings {
		settings[string(s.Name)] = s.Value
	}
	return resource.Set(c.Name, settings)
}
func resolveClustersStatistics(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster, ok := resource.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("expected to have types.Cluster but got %T", resource.Item)
	}
	stats := make(map[string]*string)
	for _, s := range cluster.Statistics {
		stats[*s.Name] = s.Value
	}
	return resource.Set(c.Name, stats)
}

func resolveClusterAttachmentsDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	attachment, ok := resource.Item.(types.Attachment)
	if !ok {
		return fmt.Errorf("expected to have types.Attachment but got %T", resource.Item)
	}
	details := make(map[string]*string)
	for _, s := range attachment.Details {
		details[*s.Name] = s.Value
	}
	return resource.Set(c.Name, details)
}
func fetchEcsClusterTasks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("expected to have types.Cluster but got %T", parent.Item)
	}
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	config := ecs.ListTasksInput{
		Cluster: cluster.ClusterArn,
	}
	for {
		listTasks, err := svc.ListTasks(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
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
		describeTasks, err := svc.DescribeTasks(ctx, &describeServicesInput, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}

		res <- describeTasks.Tasks

		if listTasks.NextToken == nil {
			break
		}
		config.NextToken = listTasks.NextToken
	}
	return nil
}

func resolveClusterTaskAttachmentsDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.Attachment)
	j := make(map[string]interface{})
	for _, i := range p.Details {
		j[*i.Name] = *i.Value
	}

	return resource.Set(c.Name, j)
}

func fetchEcsClusterServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster := parent.Item.(types.Cluster)
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	config := ecs.ListServicesInput{
		Cluster: cluster.ClusterArn,
	}
	for {
		listServicesOutput, err := svc.ListServices(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		if len(listServicesOutput.ServiceArns) == 0 {
			return nil
		}
		describeServicesInput := ecs.DescribeServicesInput{
			Cluster:  cluster.ClusterArn,
			Services: listServicesOutput.ServiceArns,
			Include:  []types.ServiceField{types.ServiceFieldTags},
		}
		describeServicesOutput, err := svc.DescribeServices(ctx, &describeServicesInput, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}

		res <- describeServicesOutput.Services

		if listServicesOutput.NextToken == nil {
			break
		}
		config.NextToken = listServicesOutput.NextToken
	}
	return nil
}

func resolveClusterServicesPlacementConstraints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service := resource.Item.(types.Service)
	j := make(map[string]interface{})
	for _, i := range service.PlacementConstraints {
		j[string(i.Type)] = aws.ToString(i.Expression)
	}

	return resource.Set(c.Name, j)
}

func resolveClusterServicesPlacementStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service := resource.Item.(types.Service)
	j := make(map[string]interface{})
	for _, i := range service.PlacementStrategy {
		j[string(i.Type)] = aws.ToString(i.Field)
	}

	return resource.Set(c.Name, j)
}

func fetchEcsClusterContainerInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster := parent.Item.(types.Cluster)
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	config := ecs.ListContainerInstancesInput{
		Cluster: cluster.ClusterArn,
	}
	for {
		listContainerInstances, err := svc.ListContainerInstances(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		if len(listContainerInstances.ContainerInstanceArns) == 0 {
			return nil
		}
		describeServicesInput := ecs.DescribeContainerInstancesInput{
			Cluster:            cluster.ClusterArn,
			ContainerInstances: listContainerInstances.ContainerInstanceArns,
			Include:            []types.ContainerInstanceField{types.ContainerInstanceFieldTags},
		}
		describeContainerInstances, err := svc.DescribeContainerInstances(ctx, &describeServicesInput, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}

		res <- describeContainerInstances.ContainerInstances

		if listContainerInstances.NextToken == nil {
			break
		}
		config.NextToken = listContainerInstances.NextToken
	}
	return nil
}

func resolveClusterContainerInstanceAttachmentsDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	attachment := resource.Item.(types.Attachment)
	details := make(map[string]*string)
	for _, s := range attachment.Details {
		details[*s.Name] = s.Value
	}
	return resource.Set(c.Name, details)
}
