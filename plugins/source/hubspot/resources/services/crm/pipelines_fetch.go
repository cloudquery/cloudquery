package crm

import (
	"context"

	"github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/pipelines"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchPipelines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	hubspotClient := pipelines.NewAPIClient(pipelines.NewConfiguration())
	cqClient := meta.(*client.Client)

	if err := cqClient.RateLimiter.Wait(ctx); err != nil {
		return nil
	}

	out, _, err := hubspotClient.PipelinesApi.GetAll(
		hubspot.WithAuthorizer(ctx, cqClient.Authorizer),
		cqClient.ObjectType).Execute()
	if err != nil {
		return err
	}

	res <- out.Results

	return nil
}
