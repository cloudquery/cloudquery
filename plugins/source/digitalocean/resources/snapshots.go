package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_snapshots",
		Description:  "Snapshot represents a DigitalOcean Snapshot",
		Resolver:     fetchSnapshots,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the snapshot.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "A human-readable name for the snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_id",
				Description: "The unique identifier for the resource that the snapshot originated from.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceID"),
			},
			{
				Name:        "resource_type",
				Description: "The type of resource that the snapshot originated from.",
				Type:        schema.TypeString,
			},
			{
				Name:        "regions",
				Description: "An array of the regions that the snapshot is available in. The regions are represented by their identifying slug values.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "min_disk_size",
				Description: "The minimum size in GB required for a volume or Droplet to use this snapshot.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "size_giga_bytes",
				Description: "The billable size of the snapshot in gigabytes.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "created",
				Description: "A time value given in ISO8601 combined date and time format that represents when the snapshot was created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "An array of Tags the snapshot has been tagged with.",
				Type:        schema.TypeStringArray,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		snapshots, resp, err := svc.DoClient.Snapshots.List(ctx, opt)
		if err != nil {
			return diag.WrapError(err)
		}
		// pass the current page's project to our result channel
		res <- snapshots
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
