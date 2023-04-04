package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	tableName := "aws_dax_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_dax_Cluster.html`,
		Resolver:    fetchDaxClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "dax"),
		Transform:   transformers.TransformWithStruct(&types.Cluster{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveClusterTags,
			},
		},
	}
}

func fetchDaxClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Dax

	config := dax.DescribeClustersInput{}
	for {
		output, err := svc.DescribeClusters(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.Clusters

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
func resolveClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster := resource.Item.(types.Cluster)

	cl := meta.(*client.Client)
	svc := cl.Services().Dax
	// TODO: Support Pagination
	response, err := svc.ListTags(ctx, &dax.ListTagsInput{
		ResourceName: cluster.ClusterArn,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(response.Tags))
}
