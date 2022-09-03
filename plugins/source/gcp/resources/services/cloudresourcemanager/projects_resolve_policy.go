package cloudresourcemanager

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/cloudresourcemanager/v3"
)

func resolveProjectPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	p := resource.Item.(*cloudresourcemanager.Project)
	output, err := cl.Services.Cloudresourcemanager.Projects.
		GetIamPolicy("projects/"+p.ProjectId, &cloudresourcemanager.GetIamPolicyRequest{}).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	var policy map[string]interface{}
	data, err := json.Marshal(output)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &policy); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, policy))
}
