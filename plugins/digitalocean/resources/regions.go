package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Regions() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_regions",
		Description:  "Region represents a DigitalOcean Region",
		Resolver:     fetchRegions,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"slug"}},
		Columns: []schema.Column{
			{
				Name:        "slug",
				Description: "A human-readable string that is used as a unique identifier for each region.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The display name of the region.  This will be a full name that is used in the control panel and other interfaces.",
				Type:        schema.TypeString,
			},
			{
				Name:        "sizes",
				Description: "This attribute is set to an array which contains the identifying slugs for the sizes available in this region.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "available",
				Description: "This is a boolean value that represents whether new Droplets can be created in this region.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "features",
				Description: "This attribute is set to an array which contains features available in this region",
				Type:        schema.TypeStringArray,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchRegions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		regions, resp, err := svc.DoClient.Regions.List(ctx, opt)
		if err != nil {
			return diag.WrapError(err)
		}
		// pass the current page's project to our result channel
		res <- regions
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return diag.WrapError(err)
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}
