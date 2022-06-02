package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

type floatingIpWrapper struct {
	godo.FloatingIP
	DropletId int
}

func FloatingIps() *schema.Table {
	return &schema.Table{
		Name:          "digitalocean_floating_ips",
		Description:   "FloatingIP represents a Digital Ocean floating IP.",
		Resolver:      fetchFloatingIps,
		DeleteFilter:  client.DeleteFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"ip"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "droplet_id",
				Description: "Unique identifier of Droplet assigned the floating ip.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "region_slug",
				Description: "A human-readable string that is used as a unique identifier for each region.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Region.Slug"),
			},
			{
				Name:        "region_name",
				Description: "The display name of the region.  This will be a full name that is used in the control panel and other interfaces.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Region.Name"),
			},
			{
				Name:        "region_sizes",
				Description: "This attribute is set to an array which contains the identifying slugs for the sizes available in this region.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Region.Sizes"),
			},
			{
				Name:        "region_available",
				Description: "This is a boolean value that represents whether new Droplets can be created in this region.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Region.Available"),
			},
			{
				Name:        "region_features",
				Description: "This attribute is set to an array which contains features available in this region",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Region.Features"),
			},
			{
				Name:        "ip",
				Description: "The public IP address of the floating IP. It also serves as its identifier.",
				Type:        schema.TypeCIDR,
				Resolver:    schema.IPNetResolver("IP"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFloatingIps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		floatIps, resp, err := svc.DoClient.FloatingIPs.List(ctx, opt)
		if err != nil {
			return diag.WrapError(err)
		}
		// pass the current page's project to our result channel
		fw := make([]floatingIpWrapper, len(floatIps))
		for i, f := range floatIps {
			fw[i] = floatingIpWrapper{
				FloatingIP: f,
				DropletId:  f.Droplet.ID,
			}
		}
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
