// Code generated by codegen; DO NOT EDIT.

package price_rule

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func PriceRules() *schema.Table {
	return &schema.Table{
		Name:     "shopify_price_rules",
		Resolver: fetchPriceRules,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "value_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ValueType"),
			},
			{
				Name:     "value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Value"),
			},
			{
				Name:     "customer_selection",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerSelection"),
			},
			{
				Name:     "target_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetType"),
			},
			{
				Name:     "target_selection",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetSelection"),
			},
			{
				Name:     "allocation_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AllocationMethod"),
			},
			{
				Name:     "once_per_customer",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("OncePerCustomer"),
			},
			{
				Name:     "usage_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("UsageLimit"),
			},
			{
				Name:     "starts_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartsAt"),
			},
			{
				Name:     "ends_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EndsAt"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "entitled_product_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EntitledProductIds"),
			},
			{
				Name:     "entitled_variant_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EntitledVariantIds"),
			},
			{
				Name:     "entitled_collection_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EntitledCollectionIds"),
			},
			{
				Name:     "entitled_country_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EntitledCountryIds"),
			},
			{
				Name:     "prerequisite_product_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrerequisiteProductIds"),
			},
			{
				Name:     "prerequisite_variant_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrerequisiteVariantIds"),
			},
			{
				Name:     "prerequisite_collection_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrerequisiteCollectionIds"),
			},
			{
				Name:     "customer_segment_prerequisite_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CustomerSegmentPrerequisiteIds"),
			},
			{
				Name:     "prerequisite_customer_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrerequisiteCustomerIds"),
			},
			{
				Name:     "prerequisite_to_entitlement_quantity_ratio",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrerequisiteToEntitlementQuantityRatio"),
			},
			{
				Name:     "prerequisite_to_entitlement_purchase",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrerequisiteToEntitlementPurchase"),
			},
			{
				Name:     "title",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Title"),
			},
			{
				Name:     "admin_graphql_api_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdminGraphqlAPIID"),
			},
		},

		Relations: []*schema.Table{
			PriceRuleDiscountCodes(),
		},
	}
}
