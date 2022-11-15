package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIamUserServicesLastAccessed(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*types.User)
	c := meta.(*client.Client)
	svc := c.Services().Iam
	return fetchIamAccessDetails(ctx, res, svc, *p.Arn)
}
