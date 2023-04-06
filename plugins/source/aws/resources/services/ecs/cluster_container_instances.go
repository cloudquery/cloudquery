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

func clusterContainerInstances() *schema.Table {
	tableName := "aws_ecs_cluster_container_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerInstance.html`,
		Resolver:    fetchEcsClusterContainerInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ecs"),
		Transform:   transformers.TransformWithStruct(&types.ContainerInstance{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "cluster_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEcsClusterContainerInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(types.Cluster)
	svc := meta.(*client.Client).Services().Ecs
	config := ecs.ListContainerInstancesInput{
		Cluster: cluster.ClusterArn,
	}
	for {
		listContainerInstances, err := svc.ListContainerInstances(ctx, &config)
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
		describeContainerInstances, err := svc.DescribeContainerInstances(ctx, &describeServicesInput)
		if err != nil {
			return err
		}

		res <- describeContainerInstances.ContainerInstances

		if aws.ToString(listContainerInstances.NextToken) == "" {
			break
		}
		config.NextToken = listContainerInstances.NextToken
	}
	return nil
}
