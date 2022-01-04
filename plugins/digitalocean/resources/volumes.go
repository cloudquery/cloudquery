package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Volumes() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_volumes",
		Resolver:     fetchVolumes,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the block storage volume.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
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
				Name:        "name",
				Description: "A human-readable name for the block storage volume. Must be lowercase and be composed only of numbers, letters and \"-\", up to a limit of 64 characters. The name must begin with a letter.",
				Type:        schema.TypeString,
			},
			{
				Name:        "size_giga_bytes",
				Description: "The size of the block storage volume in GiB (1024^3).",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "description",
				Description: "An optional free-form text field to describe a block storage volume.",
				Type:        schema.TypeString,
			},
			{
				Name:        "droplet_ids",
				Description: "An array containing the IDs of the Droplets the volume is attached to. Note that at this time, a volume can only be attached to a single Droplet.",
				Type:        schema.TypeIntArray,
				Resolver:    schema.PathResolver("DropletIDs"),
			},
			{
				Name:        "created_at",
				Description: "A time value given in ISO8601 combined date and time format that represents when the block storage volume was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "filesystem_type",
				Description: "The type of filesystem currently in-use on the volume.",
				Type:        schema.TypeString,
			},
			{
				Name:        "filesystem_label",
				Description: "The label currently applied to the filesystem.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.",
				Type:        schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "digitalocean_volume_droplets",
				Description: "Droplets that are co-located on the same physical hardware",
				Resolver:    fetchVolumeDroplets,
				Columns: []schema.Column{
					{
						Name:        "volume_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_volumes table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "droplet_id",
						Description: "Unique identifier of Droplet the volume is attached to.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "volume_id",
						Description: "The unique identifier for the block storage volume.",
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

func fetchVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	opt := &godo.ListVolumeParams{
		ListOptions: &godo.ListOptions{PerPage: client.MaxItemsPerPage},
	}
	for {
		volumes, resp, err := svc.DoClient.Storage.ListVolumes(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- volumes
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.ListOptions.Page = page + 1
	}
	return nil
}
func fetchVolumeDroplets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	volume := parent.Item.(godo.Volume)
	if volume.DropletIDs == nil {
		return nil
	}
	vd := make([]volumeDroplets, len(volume.DropletIDs))
	for i, dropletId := range volume.DropletIDs {
		vd[i] = volumeDroplets{
			DropletId: dropletId,
			VolumeId:  volume.ID,
		}
	}
	res <- vd
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

type volumeDroplets struct {
	DropletId int
	VolumeId  string
}
