package iam

import (
	"context"
	"golang.org/x/sync/errgroup"

	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"

	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIamUserLastAccessedServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Iam
	ch := make(chan interface{})

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer close(ch)
		return fetchIamUsers(ctx, meta, nil, ch)
	})

	g.Go(func() error {
		for i := range ch {
			users := i.([]types.User)
			for _, user := range users {
				if err := fetchIamAccessDetails(ctx, res, svc, *user.Arn); err != nil {
					return err
				}
			}
		}
		return nil
	})

	return g.Wait()
}

func userLastAccessedServicesPreResourceResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	p := resource.Item.(models.ServiceLastAccessedEntitiesWrapper)
	c := meta.(*client.Client)
	svc := c.Services().Iam
	entities, err := fetchDetailEntities(ctx, svc, p)
	if err != nil {
		return err
	}
	p.Entities = entities
	resource.Item = p
	return nil
}
