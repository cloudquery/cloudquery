package iam

import (
	"context"
	"fmt"

	iamadmin "cloud.google.com/go/iam/admin/apiv1"
	iampb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchRoles(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""

	iamClient, err := iamadmin.NewIamClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	iamClient.CallOptions = &iamadmin.IamCallOptions{}

	for {
		req := &iampb.ListRolesRequest{
			PageSize:  1000,
			PageToken: nextPageToken,
			Parent:    fmt.Sprintf("projects/%s", c.ProjectId),
			View:      iampb.RoleView_FULL,
		}
		resp, err := iamClient.ListRoles(ctx, req, c.CallOptions...)
		if err != nil {
			return err
		}
		res <- resp.Roles

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return nil
}
