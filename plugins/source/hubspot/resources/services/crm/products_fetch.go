package crm

import (
	"context"

	"github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/products"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	if cqClient.IsIncrementalSync() {
		return syncIncrementally(ctx, cqClient, res, "products", "hs_lastmodifieddate")
	}

	return syncAllProducts(ctx, cqClient, res)
}

func syncAllProducts(ctx context.Context, cqClient *client.Client, res chan<- any) error {
	hubspotClient := products.NewAPIClient(products.NewConfiguration())

	var after string
	for {
		if err := cqClient.RateLimiter.Wait(ctx); err != nil {
			return nil
		}

		req := hubspotClient.BasicApi.
			GetPage(hubspot.WithAuthorizer(ctx, cqClient.Authorizer)).
			Properties(cqClient.Spec.TableOptions.ForTable("hubspot_crm_products").GetProperties()).
			Associations(cqClient.Spec.TableOptions.ForTable("hubspot_crm_products").GetAssociations()).
			Limit(maxPageSize)

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
		after = next.After
		if after == "" {
			break
		}
	}

	return nil
}
