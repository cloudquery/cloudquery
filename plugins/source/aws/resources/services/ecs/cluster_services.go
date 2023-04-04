package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func clusterServices() *schema.Table {
	tableName := "aws_ecs_cluster_services"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Service.html`,
		Resolver:    fetchEcsClusterServices,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ecs"),
		Transform:   transformers.TransformWithStruct(&types.Service{}, transformers.WithPrimaryKeys("ClusterArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceArn"),
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
		Relations: []*schema.Table{
			clusterTaskSets(),
		},
	}
}

func fetchEcsClusterServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(types.Cluster)
	svc := meta.(*client.Client).Services().Ecs
	config := ecs.ListServicesInput{
		Cluster: cluster.ClusterArn,
	}
	for {
		listServicesOutput, err := svc.ListServices(ctx, &config)
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
		describeServicesOutput, err := svc.DescribeServices(ctx, &describeServicesInput)
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
