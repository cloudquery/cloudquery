// Code generated by codegen; DO NOT EDIT.

package elasticache

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ReservedCacheNodes() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_reserved_cache_nodes",
		Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReservedCacheNode.html",
		Resolver:    fetchElasticacheReservedCacheNodes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReservationARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "cache_node_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CacheNodeCount"),
			},
			{
				Name:     "cache_node_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheNodeType"),
			},
			{
				Name:     "duration",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Duration"),
			},
			{
				Name:     "fixed_price",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("FixedPrice"),
			},
			{
				Name:     "offering_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OfferingType"),
			},
			{
				Name:     "product_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProductDescription"),
			},
			{
				Name:     "recurring_charges",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RecurringCharges"),
			},
			{
				Name:     "reserved_cache_node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReservedCacheNodeId"),
			},
			{
				Name:     "reserved_cache_nodes_offering_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReservedCacheNodesOfferingId"),
			},
			{
				Name:     "start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartTime"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "usage_price",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("UsagePrice"),
			},
		},
	}
}
