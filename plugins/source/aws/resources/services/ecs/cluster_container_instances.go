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
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEcsClusterContainerInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(types.Cluster)
	cl := meta.(*client.Client)
	svc := cl.Services().Ecs
	config := ecs.ListContainerInstancesInput{
		Cluster: cluster.ClusterArn,
	}
	paginator := ecs.NewListContainerInstancesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ecs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.ContainerInstanceArns) == 0 {
			continue
		}
		describeServicesInput := ecs.DescribeContainerInstancesInput{
			Cluster:            cluster.ClusterArn,
			ContainerInstances: page.ContainerInstanceArns,
			Include:            []types.ContainerInstanceField{types.ContainerInstanceFieldTags},
		}
		describeContainerInstances, err := svc.DescribeContainerInstances(ctx, &describeServicesInput, func(options *ecs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- describeContainerInstances.ContainerInstances
	}
	return nil
}
