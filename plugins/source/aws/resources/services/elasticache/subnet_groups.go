package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SubnetGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_subnet_groups",
		Description: "Contains information about cache subnet groups",
		Resolver:    fetchElasticacheSubnetGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:            "arn",
				Description:     "The ARN (Amazon Resource Name) of the cache subnet group.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "cache_subnet_group_description",
				Description: "The description of the cache subnet group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cache_subnet_group_name",
				Description: "The name of the cache subnet group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_id",
				Description: "The Amazon Virtual Private Cloud identifier (VPC ID) of the cache subnet group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnets",
				Description: "Represents the subnet associated with a cluster",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Subnets"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchElasticacheSubnetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := elasticache.NewDescribeCacheSubnetGroupsPaginator(meta.(*client.Client).Services().ElastiCache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.CacheSubnetGroups
	}
	return nil
}
