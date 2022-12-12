package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/iam/v1"
)

func fetchServiceAccountKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	p := parent.Item.(*iam.ServiceAccount)
	iamClient, err := iam.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	output, err := iamClient.Projects.ServiceAccounts.Keys.List(p.Name).Context(ctx).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	res <- output.Keys
	return nil
}
