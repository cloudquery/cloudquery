package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ReservedInstances() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_reserved_instances",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ReservedDBInstance.html`,
		Resolver:    fetchRdsReservedInstances,
		Transform:   transformers.TransformWithStruct(&types.ReservedDBInstance{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReservedDBInstanceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsRITags,
			},
		},
	}
}

func fetchRdsReservedInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config rds.DescribeReservedDBInstancesInput
	c := meta.(*client.Client)
	svc := c.Services().Rds
	paginator := rds.NewDescribeReservedDBInstancesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.ReservedDBInstances
	}
	return nil
}

func resolveRdsRITags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Rds
	out, err := svc.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{ResourceName: resource.Item.(types.ReservedDBInstance).ReservedDBInstanceArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
