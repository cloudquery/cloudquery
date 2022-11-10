package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
)

func fetchIamPolicyServicesLastAccessed(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Iam
	ch := make(chan interface{})

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := fetchIamPolicies(ctx, meta, nil, ch); err != nil {
			return err
		}
		close(ch)
		return nil
	})

	g.Go(func() error {
		for i := range ch {
			policies := i.([]types.ManagedPolicyDetail)
			for _, r := range policies {
				if err := fetchIamAccessDetails(ctx, res, svc, *r.Arn); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return g.Wait()
}
