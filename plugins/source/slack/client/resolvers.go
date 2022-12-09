package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func ResolveTeamID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	client := meta.(*Client)
	return r.Set(col.Name, client.TeamID)
}
