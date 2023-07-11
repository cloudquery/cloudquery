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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ServiceArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
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
	cl := meta.(*client.Client)
	svc := cl.Services().Ecs
	config := ecs.ListServicesInput{
		Cluster: cluster.ClusterArn,
	}
	paginator := ecs.NewListServicesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ecs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.ServiceArns) == 0 {
			continue
		}
		describeServicesInput := ecs.DescribeServicesInput{
			Cluster:  cluster.ClusterArn,
			Services: page.ServiceArns,
			Include:  []types.ServiceField{types.ServiceFieldTags},
		}
		describeServicesOutput, err := svc.DescribeServices(ctx, &describeServicesInput, func(options *ecs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- describeServicesOutput.Services
	}
	return nil
}
