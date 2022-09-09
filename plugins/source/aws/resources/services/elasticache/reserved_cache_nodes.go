package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ReservedCacheNodes() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_reserved_cache_nodes",
		Description: "Reserved Elasticache Cache Nodes",
		Resolver:    fetchElasticacheReservedCacheNodes,
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
				Name:        "cache_node_count",
				Description: "The number of cache nodes that have been reserved.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "cache_node_type",
				Description: "The cache node type for the reserved cache nodes",
				Type:        schema.TypeString,
			},
			{
				Name:        "duration",
				Description: "The duration of the reservation in seconds.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "fixed_price",
				Description: "The fixed price charged for this reserved cache node.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "offering_type",
				Description: "The offering type of this reserved cache node.",
				Type:        schema.TypeString,
			},
			{
				Name:        "product_description",
				Description: "The description of the reserved cache node.",
				Type:        schema.TypeString,
			},
			{
				Name:            "reservation_arn",
				Description:     "The Amazon Resource Name (ARN) of the reserved cache node",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ReservationARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "reserved_cache_node_id",
				Description: "The unique identifier for the reservation.",
				Type:        schema.TypeString,
			},
			{
				Name:        "reserved_cache_nodes_offering_id",
				Description: "The offering identifier.",
				Type:        schema.TypeString,
			},
			{
				Name:        "start_time",
				Description: "The time the reservation started.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "state",
				Description: "The state of the reserved cache node.",
				Type:        schema.TypeString,
			},
			{
				Name:        "usage_price",
				Description: "The hourly price charged for this reserved cache node.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "recurring_charges",
				Description: "Contains the specific price and frequency of a recurring charges for a reserved cache node, or for a reserved cache node offering.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("RecurringCharges"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchElasticacheReservedCacheNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := elasticache.NewDescribeReservedCacheNodesPaginator(meta.(*client.Client).Services().ElastiCache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.ReservedCacheNodes
	}
	return nil
}
