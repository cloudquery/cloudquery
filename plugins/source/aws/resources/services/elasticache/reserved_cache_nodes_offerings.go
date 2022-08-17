package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource reserved_cache_nodes_offerings --config ./gen.hcl --output .
func ReservedCacheNodesOfferings() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticache_reserved_cache_nodes_offerings",
		Description:  "Describes all of the attributes of a reserved cache node offering.",
		Resolver:     fetchElasticacheReservedCacheNodesOfferings,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticache"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "reserved_cache_nodes_offering_id"}},
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
				Name:        "cache_node_type",
				Description: "The cache node type for the reserved cache node",
				Type:        schema.TypeString,
			},
			{
				Name:        "duration",
				Description: "The duration of the offering",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "fixed_price",
				Description: "The fixed price charged for this offering.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "offering_type",
				Description: "The offering type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "product_description",
				Description: "The cache engine used by the offering.",
				Type:        schema.TypeString,
			},
			{
				Name:        "reserved_cache_nodes_offering_id",
				Description: "A unique identifier for the reserved cache node offering.",
				Type:        schema.TypeString,
			},
			{
				Name:        "usage_price",
				Description: "The hourly price charged for this offering.",
				Type:        schema.TypeFloat,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elasticache_reserved_cache_nodes_offering_recurring_charges",
				Description: "Contains the specific price and frequency of a recurring charges for a reserved cache node, or for a reserved cache node offering.",
				Resolver:    schema.PathTableResolver("RecurringCharges"),
				Columns: []schema.Column{
					{
						Name:        "reserved_cache_nodes_offering_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_reserved_cache_nodes_offerings table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "recurring_charge_amount",
						Description: "The monetary amount of the recurring charge.",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "recurring_charge_frequency",
						Description: "The frequency of the recurring charge.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchElasticacheReservedCacheNodesOfferings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := elasticache.NewDescribeReservedCacheNodesOfferingsPaginator(meta.(*client.Client).Services().ElastiCache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.ReservedCacheNodesOfferings
	}
	return nil
}
