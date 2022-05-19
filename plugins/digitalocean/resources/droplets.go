package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

type neighborWrapper struct {
	DropletId  int
	NeighborId int
}

func Droplets() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_droplets",
		Description:  "Droplet represents a DigitalOcean Droplet",
		Resolver:     fetchDroplets,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "A unique identifier for each Droplet instance. This is automatically generated upon Droplet creation.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The human-readable name set for the Droplet instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "memory",
				Description: "Memory of the Droplet in megabytes.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "vcpus",
				Description: "The number of virtual CPUs.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "disk",
				Description: "The size of the Droplet's disk in gigabytes.",
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
				Name:        "image_id",
				Description: "A unique number that can be used to identify and reference a specific image.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Image.ID"),
			},
			{
				Name:        "image_name",
				Description: "The display name that has been given to an image.  This is what is shown in the control panel and is generally a descriptive title for the image in question.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Image.Name"),
			},
			{
				Name:        "image_type",
				Description: "Describes the kind of image. It may be one of \"snapshot\", \"backup\", or \"custom\". This specifies whether an image is a user-generated Droplet snapshot, automatically created Droplet backup, or a user-provided virtual machine image.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Image.Type"),
			},
			{
				Name:        "image_distribution",
				Description: "The name of a custom image's distribution. Currently, the valid values are  \"Arch Linux\", \"CentOS\", \"CoreOS\", \"Debian\", \"Fedora\", \"Fedora Atomic\",  \"FreeBSD\", \"Gentoo\", \"openSUSE\", \"RancherOS\", \"Ubuntu\", and \"Unknown\".  Any other value will be accepted but ignored, and \"Unknown\" will be used in its place.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Image.Distribution"),
			},
			{
				Name:        "image_slug",
				Description: "A uniquely identifying string that is associated with each of the DigitalOcean-provided public images. These can be used to reference a public image as an alternative to the numeric id.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Image.Slug"),
			},
			{
				Name:        "image_public",
				Description: "This is a boolean value that indicates whether the image in question is public or not. An image that is public is available to all accounts. A non-public image is only accessible from your account.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Image.Public"),
			},
			{
				Name:        "image_regions",
				Description: "This attribute is an array of the regions that the image is available in. The regions are represented by their identifying slug values.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Image.Regions"),
			},
			{
				Name:        "image_min_disk_size",
				Description: "The minimum disk size in GB required for a Droplet to use this image.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Image.MinDiskSize"),
			},
			{
				Name:        "image_size_giga_bytes",
				Description: "The size of the image in gigabytes.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Image.SizeGigaBytes"),
			},
			{
				Name:        "image_created",
				Description: "A time value given in ISO8601 combined date and time format that represents when the image was created.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Image.Created"),
			},
			{
				Name:        "image_description",
				Description: "An optional free-form text field to describe an image.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Image.Description"),
			},
			{
				Name:        "image_tags",
				Description: "A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Image.Tags"),
			},
			{
				Name:        "image_status",
				Description: "A status string indicating the state of a custom image. This may be `NEW`,\n `available`, `pending`, `deleted`, or `retired`.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Image.Status"),
			},
			{
				Name:        "image_error_message",
				Description: "A string containing information about errors that may occur when importing\n a custom image.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Image.ErrorMessage"),
			},
			{
				Name:        "size_memory",
				Description: "The amount of RAM allocated to Droplets created of this size. The value is represented in megabytes.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Size.Memory"),
			},
			{
				Name:     "size_vcpus",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Size.Vcpus"),
			},
			{
				Name:        "size_disk",
				Description: "The amount of disk space set aside for Droplets of this size. The value is represented in gigabytes.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Size.Disk"),
			},
			{
				Name:     "size_price_monthly",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Size.PriceMonthly"),
			},
			{
				Name:        "size_price_hourly",
				Description: "This describes the price of the Droplet size as measured hourly. The value is measured in US dollars.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Size.PriceHourly"),
			},
			{
				Name:     "size_regions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Size.Regions"),
			},
			{
				Name:        "size_available",
				Description: "This is a boolean value that represents whether new Droplets can be created with this size.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Size.Available"),
			},
			{
				Name:        "size_transfer",
				Description: "The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Size.Transfer"),
			},
			{
				Name:     "size_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Size.Description"),
			},
			{
				Name:        "size_slug",
				Description: "The unique slug identifier for the size of this Droplet.",
				Type:        schema.TypeString,
			},
			{
				Name:        "backup_ids",
				Description: "An array of backup IDs of any backups that have been taken of the Droplet instance.  Droplet backups are enabled at the time of the instance creation.",
				Type:        schema.TypeIntArray,
				Resolver:    schema.PathResolver("BackupIDs"),
			},
			{
				Name:        "next_backup_window_start_time",
				Description: "A time value given in ISO8601 combined date and time format specifying the start of the Droplet's backup window.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("NextBackupWindow.Start.Time"),
			},
			{
				Name:        "next_backup_window_end_time",
				Description: "A time value given in ISO8601 combined date and time format specifying the end of the Droplet's backup window.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("NextBackupWindow.End.Time"),
			},
			{
				Name:        "snapshot_ids",
				Description: "An array of snapshot IDs of any snapshots created from the Droplet instance.",
				Type:        schema.TypeIntArray,
				Resolver:    schema.PathResolver("SnapshotIDs"),
			},
			{
				Name:        "features",
				Description: "An array of features enabled on this Droplet.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "locked",
				Description: "A boolean value indicating whether the Droplet has been locked, preventing actions by users.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "status",
				Description: "A status string indicating the state of the Droplet instance. This may be \"new\", \"active\", \"off\", or \"archive\".",
				Type:        schema.TypeString,
			},
			{
				Name:        "created",
				Description: "A time value given in ISO8601 combined date and time format that represents when the Droplet was created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "kernel_id",
				Description: "A unique number used to identify and reference a specific kernel.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Kernel.ID"),
			},
			{
				Name:        "kernel_name",
				Description: "The display name of the kernel. This is shown in the web UI and is generally a descriptive title for the kernel in question.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Kernel.Name"),
			},
			{
				Name:        "kernel_version",
				Description: "A standard kernel version string representing the version, patch, and release information.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Kernel.Version"),
			},
			{
				Name:        "tags",
				Description: "An array of Tags the Droplet has been tagged with.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "volume_ids",
				Description: "A flat array including the unique identifier for each Block Storage volume attached to the Droplet.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VolumeIDs"),
			},
			{
				Name:        "vpc_uuid",
				Description: "A string specifying the UUID of the VPC to which the Droplet is assigned.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VPCUUID"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "digitalocean_droplet_networks_v4",
				Description: "NetworkV4 represents a DigitalOcean IPv4 Network.",
				Resolver:    fetchDropletNetworksV4,
				Columns: []schema.Column{
					{
						Name:        "droplet_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_droplets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "ip_address",
						Description: "The IP address of the IPv4 network interface.",
						Type:        schema.TypeInet,
						Resolver:    client.IPAddressResolver("IPAddress"),
					},
					{
						Name:        "netmask",
						Description: "The netmask of the IPv4 network interface.",
						Type:        schema.TypeInet,
						Resolver:    client.IPAddressResolver("Netmask"),
					},
					{
						Name:        "gateway",
						Description: "The gateway of the specified IPv4 network interface.\n\nFor private interfaces, a gateway is not provided. This is denoted by\nreturning `nil` as its value.\n",
						Type:        schema.TypeInet,
						Resolver:    client.IPAddressResolver("Gateway"),
					},
					{
						Name:        "type",
						Description: "The type of the IPv4 network interface.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "digitalocean_droplet_networks_v6",
				Description: "NetworkV6 represents a DigitalOcean IPv6 network.",
				Resolver:    fetchDropletNetworksV6,
				Columns: []schema.Column{
					{
						Name:        "droplet_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_droplets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "ip_address",
						Description: "The IP address of the IPv6 network interface.",
						Type:        schema.TypeInet,
						Resolver:    client.IPAddressResolver("IPAddress"),
					},
					{
						Name:        "netmask",
						Description: "The netmask of the IPv6 network interface.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Netmask"),
					},
					{
						Name:        "gateway",
						Description: "The gateway of the specified IPv6 network interface.",
						Type:        schema.TypeInet,
						Resolver:    client.IPAddressResolver("Gateway"),
					},
					{
						Name:        "type",
						Description: "The type of the IPv6 network interface.\n\n**Note**: IPv6 private  networking is not currently supported.\n",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "digitalocean_droplet_neighbors",
				Description:   "Droplets that are co-located on the same physical hardware",
				Resolver:      fetchDropletNeighbors,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "droplet_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_droplets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "droplet_id",
						Description: "Unique identifier of the droplet associated with the neighbor droplet.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "neighbor_id",
						Description: "Droplet neighbor identifier that exists on same the same physical hardware as the droplet.",
						Type:        schema.TypeBigInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchDroplets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		droplets, resp, err := svc.DoClient.Droplets.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- droplets
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}
func fetchDropletNetworksV4(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	droplet := parent.Item.(godo.Droplet)
	if droplet.Networks == nil {
		return nil
	}
	res <- droplet.Networks.V4
	return nil
}
func fetchDropletNetworksV6(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	droplet := parent.Item.(godo.Droplet)
	if droplet.Networks == nil {
		return nil
	}
	res <- droplet.Networks.V6
	return nil
}
func fetchDropletNeighbors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	droplet := parent.Item.(godo.Droplet)

	neighbors, _, err := svc.DoClient.Droplets.Neighbors(ctx, droplet.ID)
	if err != nil {
		return err
	}
	if neighbors == nil {
		return nil
	}
	nn := make([]neighborWrapper, len(neighbors))
	for i, n := range neighbors {
		nn[i] = neighborWrapper{
			DropletId:  droplet.ID,
			NeighborId: n.ID,
		}
	}
	res <- nn
	return nil
}
