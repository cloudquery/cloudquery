package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEcsClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ecs.ListClustersInput
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().Ecs
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

func fetchEcsClusterTasks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("expected to have types.Cluster but got %T", parent.Item)
	}
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().Ecs
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

func fetchEcsClusterServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster := parent.Item.(types.Cluster)
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().Ecs
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

func fetchEcsClusterContainerInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster := parent.Item.(types.Cluster)
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().Ecs
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
