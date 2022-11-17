// Code generated by codegen; DO NOT EDIT.

package security

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchJitNetworkAccessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Security

	pager := svc.JitNetworkAccessPoliciesClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Value
	}

	return nil
}
