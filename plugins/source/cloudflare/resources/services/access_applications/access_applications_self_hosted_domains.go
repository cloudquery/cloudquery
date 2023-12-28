package access_applications

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

type AccessApplicationsSelfHostedDomains struct {
	ApplicationID string `json:"application_id"`
	Domain        string `json:"domain"`
}

func SelfHostedDomainsTable() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_access_applications_self_hosted_domains",
		Transform: transformers.TransformWithStruct(&AccessApplicationsSelfHostedDomains{}, transformers.WithPrimaryKeys("ApplicationID", "Domain")),
		Resolver:  fetchSelfHostedDomains,
	}
}

func fetchSelfHostedDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	application := parent.Item.(cloudflare.AccessApplication)

	for _, domain := range application.SelfHostedDomains {
		row := AccessApplicationsSelfHostedDomains{
			ApplicationID: application.ID,
			Domain:        domain,
		}
		res <- row
	}

	return nil
}
