package client

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func ResolveAWSAccount(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	r.Set("account_id", client.AccountID)
	return nil
}

func ResolveAWSRegion(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	r.Set("region", client.Region)
	return nil
}
