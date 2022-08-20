package cloudbilling

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/cloudbilling/v1"
)

//go:generate cq-gen --resource services --config gen.hcl --output .
func Services() *schema.Table {
	return &schema.Table{
		Name:          "gcp_cloudbilling_services",
		Description:   "Encapsulates a single service in Google Cloud Platform",
		Resolver:      fetchCloudbillingServices,
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "business_entity_name",
				Description: "The business under which the service is offered Ex",
				Type:        schema.TypeString,
			},
			{
				Name:        "display_name",
				Description: "A human readable display name for this service",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The resource name for the service",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_id",
				Description: "The identifier for the service",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_cloudbilling_service_skus",
				Description: "Encapsulates a single SKU in Google Cloud Platform",
				Resolver:    fetchCloudbillingServiceSkus,

				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudbilling_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "resource_family",
						Description: "The type of product the SKU refers to",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Category.ResourceFamily"),
					},
					{
						Name:        "resource_group",
						Description: "A group classification for related SKUs",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Category.ResourceGroup"),
					},
					{
						Name:        "service_display_name",
						Description: "The display name of the service this SKU belongs to",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Category.ServiceDisplayName"),
					},
					{
						Name:        "usage_type",
						Description: "Represents how the SKU is consumed",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Category.UsageType"),
					},
					{
						Name:        "description",
						Description: "A human readable description of the SKU, has a maximum length of 256 characters",
						Type:        schema.TypeString,
					},
					{
						Name:        "geo_taxonomy_regions",
						Description: "The list of regions associated with a sku",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("GeoTaxonomy.Regions"),
					},
					{
						Name:        "geo_taxonomy_type",
						Description: "\"TYPE_UNSPECIFIED\" - The type is not specified   \"GLOBAL\" - The sku is global in nature, eg",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("GeoTaxonomy.Type"),
					},
					{
						Name:        "name",
						Description: "The resource name for the SKU",
						Type:        schema.TypeString,
					},
					{
						Name:        "service_provider_name",
						Description: "Identifies the service provider",
						Type:        schema.TypeString,
					},
					{
						Name:        "service_regions",
						Description: "\"asia-east1\" Service regions can be found at https://cloudgooglecom/about/locations/",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "sku_id",
						Description: "The identifier for the SKU",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "gcp_cloudbilling_service_sku_pricing_info",
						Description: "Represents the pricing information for a SKU at a single point of time",
						Resolver:    fetchCloudbillingServiceSkuPricingInfos,
						Columns: []schema.Column{
							{
								Name:        "service_sku_cq_id",
								Description: "Unique CloudQuery ID of gcp_cloudbilling_service_skus table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "aggregation_count",
								Description: "The number of intervals to aggregate over",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("AggregationInfo.AggregationCount"),
							},
							{
								Name:        "aggregation_interval",
								Description: "\"AGGREGATION_INTERVAL_UNSPECIFIED\"   \"DAILY\"   \"MONTHLY\"",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("AggregationInfo.AggregationInterval"),
							},
							{
								Name:        "aggregation_level",
								Description: "\"AGGREGATION_LEVEL_UNSPECIFIED\"   \"ACCOUNT\"   \"PROJECT\"",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("AggregationInfo.AggregationLevel"),
							},
							{
								Name:        "currency_conversion_rate",
								Description: "Conversion rate used for currency conversion, from USD to the currency specified in the request",
								Type:        schema.TypeFloat,
							},
							{
								Name:        "effective_time",
								Description: "The timestamp from which this pricing was effective within the requested time range",
								Type:        schema.TypeString,
							},
							{
								Name:        "base_unit",
								Description: "The base unit for the SKU which is the unit used in usage exports",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PricingExpression.BaseUnit"),
							},
							{
								Name:        "base_unit_conversion_factor",
								Description: "Conversion factor for converting from price per usage_unit to price per base_unit, and start_usage_amount to start_usage_amount in base_unit",
								Type:        schema.TypeFloat,
								Resolver:    schema.PathResolver("PricingExpression.BaseUnitConversionFactor"),
							},
							{
								Name:        "base_unit_description",
								Description: "The base unit in human readable form",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PricingExpression.BaseUnitDescription"),
							},
							{
								Name:        "display_quantity",
								Description: "The recommended quantity of units for displaying pricing info",
								Type:        schema.TypeFloat,
								Resolver:    schema.PathResolver("PricingExpression.DisplayQuantity"),
							},
							{
								Name:        "usage_unit",
								Description: "The short hand for unit of usage this pricing is specified in",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PricingExpression.UsageUnit"),
							},
							{
								Name:        "usage_unit_description",
								Description: "\"gibi byte\"",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PricingExpression.UsageUnitDescription"),
							},
							{
								Name:        "summary",
								Description: "An optional human readable summary of the pricing information, has a maximum length of 256 characters",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "gcp_cloudbilling_service_sku_pricing_info_tiered_rates",
								Description: "The price rate indicating starting usage and its corresponding price",
								Resolver:    fetchCloudbillingServiceSkuPricingInfoTieredRates,
								Columns: []schema.Column{
									{
										Name:        "service_sku_pricing_info_cq_id",
										Description: "Unique CloudQuery ID of gcp_cloudbilling_service_sku_pricing_info table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "start_usage_amount",
										Description: "Usage is priced at this rate only after this amount",
										Type:        schema.TypeFloat,
									},
									{
										Name:        "unit_price_currency_code",
										Description: "The three-letter currency code defined in ISO 4217",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("UnitPrice.CurrencyCode"),
									},
									{
										Name:        "unit_price_nanos",
										Description: "Number of nano (10^-9) units of the amount",
										Type:        schema.TypeBigInt,
										Resolver:    schema.PathResolver("UnitPrice.Nanos"),
									},
									{
										Name:        "unit_price_units",
										Description: "The whole units of the amount",
										Type:        schema.TypeBigInt,
										Resolver:    schema.PathResolver("UnitPrice.Units"),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCloudbillingServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.CloudBilling.Services.List().PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Services

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchCloudbillingServiceSkus(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*cloudbilling.Service)
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.CloudBilling.Services.Skus.List(r.Name).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Skus

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchCloudbillingServiceSkuPricingInfos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*cloudbilling.Sku)
	res <- r.PricingInfo
	return nil
}
func fetchCloudbillingServiceSkuPricingInfoTieredRates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*cloudbilling.PricingInfo)
	if r.PricingExpression == nil {
		return nil
	}
	res <- r.PricingExpression.TieredRates
	return nil
}
