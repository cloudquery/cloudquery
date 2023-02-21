package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func ResolveProjectID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cl := meta.(*Client)
	if cl.MPSpec.ProjectID < 1 {
		return nil
	}
	return r.Set(c.Name, cl.MPSpec.ProjectID)
}

func ResolveWorkspaceID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cl := meta.(*Client)
	if cl.MPSpec.WorkspaceID < 1 {
		return nil
	}
	return r.Set(c.Name, cl.MPSpec.WorkspaceID)
}
