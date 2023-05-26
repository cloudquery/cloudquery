package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v3/schema"
)

// ResolveOracleRegion grabs the region from the cq-client, not from the resource.
func ResolveOracleRegion(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("region", client.Region)
}

// ResolveCompartmentID grabs the compartment-id from the cq-client, not from the resource.
func ResolveCompartmentID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("compartment_id", client.CompartmentOcid)
}
