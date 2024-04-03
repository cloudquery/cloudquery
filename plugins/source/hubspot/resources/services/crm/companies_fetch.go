package crm

import (
	"context"
	"github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/companies"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchCompanies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	hubspotClient := companies.NewAPIClient(companies.NewConfiguration())
	cqClient := meta.(*client.Client)

	const key = "companies"

	after, err := getCursor(ctx, cqClient, key)
	if err != nil {
		return err
	}

	for {
		if err := cqClient.RateLimiter.Wait(ctx); err != nil {
			return nil
		}

		req := hubspotClient.BasicApi.
			GetPage(hubspot.WithAuthorizer(ctx, cqClient.Authorizer)).
			Properties(cqClient.Spec.TableOptions.ForTable("hubspot_crm_companies").GetProperties()).
			Associations(cqClient.Spec.TableOptions.ForTable("hubspot_crm_companies").GetAssociations()).
			Limit(client.DefaultPageSize)

		if len(after) > 0 {
			req = req.After(after)
		}
		out, _, err := req.Execute()
		if err != nil {
			return err
		}

		res <- out.Results

		if !out.HasPaging() {
			break
		}
		paging := out.GetPaging()
		if !paging.HasNext() {
			break
		}
		next := paging.GetNext()
		if next.After == "" {
			break
		}
		after = next.After
	}

	return setCursor(ctx, cqClient, key, after)
}
