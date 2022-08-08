package account

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/helpers"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource locations --config gen.hcl --output .
func Locations() *schema.Table {
	return &schema.Table{
		Name:         "azure_account_locations",
		Description:  "Azure location information",
		Resolver:     fetchAccountLocations,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "geography_group",
				Description: "The geography group of the location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.GeographyGroup"),
			},
			{
				Name:          "home_location",
				Description:   "The home location of an edge zone",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Metadata.HomeLocation"),
				IgnoreInTests: true,
			},
			{
				Name:        "latitude",
				Description: "The latitude of the location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.Latitude"),
			},
			{
				Name:        "longitude",
				Description: "The longitude of the location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.Longitude"),
			},
			{
				Name:        "physical_location",
				Description: "The physical location of the Azure location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.PhysicalLocation"),
			},
			{
				Name:        "region_category",
				Description: "The category of the region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.RegionCategory"),
			},
			{
				Name:        "region_type",
				Description: "The type of the region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.RegionType"),
			},
			{
				Name:        "display_name",
				Description: "The display name of the location",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The fully qualified ID of the location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The location name",
				Type:        schema.TypeString,
			},
			{
				Name:        "regional_display_name",
				Description: "The display name of the location and its region",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The location type",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_account_location_paired_region",
				Description: "Information regarding paired region",
				Resolver:    fetchAccountLocationPairedRegions,
				Columns: []schema.Column{
					{
						Name:        "location_cq_id",
						Description: "Unique CloudQuery ID of azure_account_locations table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The fully qualified ID of the location",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The name of the paired region",
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

func fetchAccountLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions
	pager := svc.Subscriptions.NewListLocationsPager(svc.SubscriptionID, nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, v := range nextResult.Value {
			res <- v
		}
	}
	return nil
}
func fetchAccountLocationPairedRegions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	// This is a generated base implementation. You may edit it if necessary.
	p := helpers.ToPointer(parent.Item).(*armsubscriptions.Location)
	if p.Metadata == nil {
		return nil
	}
	res <- p.Metadata.PairedRegion
	return nil
}
