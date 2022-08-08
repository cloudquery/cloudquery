package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Sizes() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_sizes",
		Description:  "Size represents a DigitalOcean Size",
		Resolver:     fetchSizes,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"slug"}},
		Columns: []schema.Column{
			{
				Name:        "slug",
				Description: "A human-readable string that is used to uniquely identify each size.",
				Type:        schema.TypeString,
			},
			{
				Name:        "memory",
				Description: "The amount of RAM allocated to Droplets created of this size. The value is represented in megabytes.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "vcpus",
				Description: "The integer of number CPUs allocated to Droplets of this size.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "disk",
				Description: "The amount of disk space set aside for Droplets of this size. The value is represented in gigabytes.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "price_monthly",
				Description: "This attribute describes the monthly cost of this Droplet size if the Droplet is kept for an entire month. The value is measured in US dollars.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "price_hourly",
				Description: "This describes the price of the Droplet size as measured hourly. The value is measured in US dollars.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "regions",
				Description: "An array containing the region slugs where this size is available for Droplet creates.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "available",
				Description: "This is a boolean value that represents whether new Droplets can be created with this size.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "transfer",
				Description: "The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "description",
				Description: "A string describing the class of Droplets created from this size. For example: Basic, General Purpose, CPU-Optimized, Memory-Optimized, or Storage-Optimized.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSizes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		sizes, resp, err := svc.DoClient.Sizes.List(ctx, opt)
		if err != nil {
			return diag.WrapError(err)
		}
		// pass the current page's project to our result channel
		res <- sizes
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
