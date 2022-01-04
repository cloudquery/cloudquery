package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:          "digitalocean_certificates",
		Description:   "Certificate represents a DigitalOcean certificate configuration.",
		Resolver:      fetchCertificates,
		DeleteFilter:  client.DeleteFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "A unique ID that can be used to identify and reference a certificate.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "A unique human-readable name referring to a certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "dns_names",
				Description: "An array of fully qualified domain names (FQDNs) for which the certificate was issued.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DNSNames"),
			},
			{
				Name:        "not_after",
				Description: "A time value given in ISO8601 combined date and time format that represents the certificate's expiration date.",
				Type:        schema.TypeString,
			},
			{
				Name:        "s_h_a1_fingerprint",
				Description: "A unique identifier generated from the SHA-1 fingerprint of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created",
				Description: "A time value given in ISO8601 combined date and time format that represents when the certificate was created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "A string representing the current state of the certificate. It may be `pending`, `verified`, or `error`.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		certs, resp, err := svc.DoClient.Certificates.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- certs
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
