package client

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ResolveOrg(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("org", client.Org)
}
