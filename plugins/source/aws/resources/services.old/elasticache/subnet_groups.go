package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource subnet_groups --config ./gen.hcl --output .
func SubnetGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticache_subnet_groups",
		Description:  "Contains information about cache subnet groups",
		Resolver:     fetchElasticacheSubnetGroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticache"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the cache subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
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
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elasticache_subnet_group_subnets",
				Description: "Represents the subnet associated with a cluster",
				Resolver:    schema.PathTableResolver("Subnets"),
				Columns: []schema.Column{
					{
						Name:        "subnet_group_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_subnet_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "subnet_availability_zone_name",
						Description: "The name of the Availability Zone.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetAvailabilityZone.Name"),
					},
					{
						Name:        "subnet_identifier",
						Description: "The unique identifier for the subnet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "subnet_outpost_arn",
						Description: "The outpost ARN of the subnet.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetOutpost.SubnetOutpostArn"),
					},
				},
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
			return diag.WrapError(err)
		}
		res <- v.CacheSubnetGroups
	}
	return nil
}
