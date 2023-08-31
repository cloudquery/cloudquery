package iam

import (
	"context"

	iamadmin "cloud.google.com/go/iam/admin/apiv1"
	iampb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func fetchServiceAccountKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := parent.Item.(*iampb.ServiceAccount)
	iamClient, err := iamadmin.NewIamClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	iamClient.CallOptions = &iamadmin.IamCallOptions{}

	req := &iampb.ListServiceAccountKeysRequest{
		Name: p.Name,
	}

	output, err := iamClient.ListServiceAccountKeys(ctx, req, c.CallOptions...)
	if err != nil {
		return err
	}

	res <- output.Keys
	return nil
}
