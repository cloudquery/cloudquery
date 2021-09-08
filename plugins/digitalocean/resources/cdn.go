package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Cdns() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_cdn",
		Resolver:     fetchCdns,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "A unique ID that can be used to identify and reference a CDN endpoint.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "origin",
				Description: "The fully qualified domain name (FQDN) for the origin server which provides the content for the CDN. This is currently restricted to a Space.",
				Type:        schema.TypeString,
			},
			{
				Name:        "endpoint",
				Description: "The fully qualified domain name (FQDN) from which the CDN-backed content is served.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "A time value given in ISO8601 combined date and time format that represents when the CDN endpoint was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "ttl",
				Description: "The amount of time the content is cached by the CDN's edge servers in seconds. TTL must be one of 60, 600, 3600, 86400, or 604800. Defaults to 3600 (one hour) when excluded.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("TTL"),
			},
			{
				Name:        "certificate_id",
				Description: "The ID of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CertificateID"),
			},
			{
				Name:        "custom_domain",
				Description: "The fully qualified domain name (FQDN) of the custom subdomain used with the CDN endpoint.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCdns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		cdns, resp, err := svc.DoClient.CDNs.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- cdns
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
