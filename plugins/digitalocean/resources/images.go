package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_images",
		Description:  "Image represents a DigitalOcean Image",
		Resolver:     fetchImages,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "A unique number that can be used to identify and reference a specific image.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The display name that has been given to an image.  This is what is shown in the control panel and is generally a descriptive title for the image in question.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Describes the kind of image. It may be one of \"snapshot\", \"backup\", or \"custom\". This specifies whether an image is a user-generated Droplet snapshot, automatically created Droplet backup, or a user-provided virtual machine image.",
				Type:        schema.TypeString,
			},
			{
				Name:        "distribution",
				Description: "The name of a custom image's distribution. Currently, the valid values are  \"Arch Linux\", \"CentOS\", \"CoreOS\", \"Debian\", \"Fedora\", \"Fedora Atomic\",  \"FreeBSD\", \"Gentoo\", \"openSUSE\", \"RancherOS\", \"Ubuntu\", and \"Unknown\".  Any other value will be accepted but ignored, and \"Unknown\" will be used in its place.",
				Type:        schema.TypeString,
			},
			{
				Name:        "slug",
				Description: "A uniquely identifying string that is associated with each of the DigitalOcean-provided public images. These can be used to reference a public image as an alternative to the numeric id.",
				Type:        schema.TypeString,
			},
			{
				Name:        "public",
				Description: "This is a boolean value that indicates whether the image in question is public or not. An image that is public is available to all accounts. A non-public image is only accessible from your account.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "regions",
				Description: "This attribute is an array of the regions that the image is available in. The regions are represented by their identifying slug values.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "min_disk_size",
				Description: "The minimum disk size in GB required for a Droplet to use this image.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "size_giga_bytes",
				Description: "The size of the image in gigabytes.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "created",
				Description: "A time value given in ISO8601 combined date and time format that represents when the image was created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional free-form text field to describe an image.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "status",
				Description: "A status string indicating the state of a custom image. This may be `NEW`,\n `available`, `pending`, `deleted`, or `retired`.",
				Type:        schema.TypeString,
			},
			{
				Name:        "error_message",
				Description: "A string containing information about errors that may occur when importing\n a custom image.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		images, resp, err := svc.DoClient.Images.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- images
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
