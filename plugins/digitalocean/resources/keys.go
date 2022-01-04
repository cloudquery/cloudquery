package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Keys() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_keys",
		Resolver:     fetchKeys,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "A unique identification number for this key. Can be used to embed specific SSH key into a Droplet.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "A human-readable display name for this key, used to easily identify the SSH keys when they are displayed.",
				Type:        schema.TypeString,
			},
			{
				Name:        "fingerprint",
				Description: "A unique identifier that differentiates this key from other keys using a format that SSH recognizes. The fingerprint is created when the key is added to your account.",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_key",
				Description: "The entire public key string that was uploaded. Embedded into the root user's `authorized_keys` file if you include this key during Droplet creation.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		keys, resp, err := svc.DoClient.Keys.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- keys
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
