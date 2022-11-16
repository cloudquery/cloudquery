package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIamPolicyLastAccessedServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(types.ManagedPolicyDetail)
	c := meta.(*client.Client)
	svc := c.Services().Iam
	return fetchIamAccessDetails(ctx, res, svc, *p.Arn)
}

func policyLastAccessedServicesPreResourceResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
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
