package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Vpcs() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_vpcs",
		Description:  "VPC represents a DigitalOcean Virtual Private Cloud configuration.",
		Resolver:     fetchVpcs,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "A unique ID that can be used to identify and reference the VPC.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "urn",
				Description: "The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("URN"),
			},
			{
				Name:        "name",
				Description: "The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "A free-form text field for describing the VPC's purpose. It may be a maximum of 255 characters.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ip_range",
				Description: "The range of IP addresses in the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be smaller than `/24` nor larger than `/16`. If no IP range is specified, a `/20` network range is generated that won't conflict with other VPC networks in your account.",
				Type:        schema.TypeCIDR,
				Resolver:    schema.IPNetResolver("IPRange"),
			},
			{
				Name:        "region_slug",
				Description: "The slug identifier for the region where the VPC will be created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "A time value given in ISO8601 combined date and time format.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "default",
				Description: "A boolean value indicating whether or not the VPC is the default network for the region. All applicable resources are placed into the default VPC network unless otherwise specified during their creation. The `default` field cannot be unset from `true`. If you want to set a new default VPC network, update the `default` field of another VPC network in the same region. The previous network's `default` field will be set to `false` when a new default VPC has been defined.",
				Type:        schema.TypeBool,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "digitalocean_vpc_members",
				Description: "Resources that are members of a VPC.",
				Resolver:    fetchVpcMembers,
				Columns: []schema.Column{
					{
						Name:        "vpc_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_vpcs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "The resource type of the URN associated with the VPC..",
						Type:        schema.TypeString,
						Resolver:    client.ResolveResourceTypeFromUrn,
					},
					{
						Name:        "id",
						Description: "A unique ID that can be used to identify the resource that is a member of the VPC.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveResourceIdFromUrn,
					},
					{
						Name:        "urn",
						Description: "The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("URN"),
					},
					{
						Name:        "name",
						Description: "The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods.",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "A time value given in ISO8601 combined date and time format.",
						Type:        schema.TypeTimestamp,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchVpcs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		vpcs, resp, err := svc.DoClient.VPCs.List(ctx, opt)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- vpcs
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
func fetchVpcMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	vpc := parent.Item.(*godo.VPC)
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		vpcMembers, resp, err := svc.DoClient.VPCs.ListMembers(ctx, vpc.ID, nil, opt)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- vpcMembers
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
