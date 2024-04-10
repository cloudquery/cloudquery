package crm

import (
	"context"
	"fmt"
	"github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/contacts"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"time"
)

func fetchContacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	hubspotClient := contacts.NewAPIClient(contacts.NewConfiguration())
	cqClient := meta.(*client.Client)

	if cqClient.IsIncrementalSync() {
		return syncContactsIncrementally(ctx, cqClient, hubspotClient, res)
	}

	return syncAllContacts(ctx, cqClient, hubspotClient, res)
}

func syncAllContacts(ctx context.Context, cqClient *client.Client, hubspotClient *contacts.APIClient, res chan<- any) error {
	var after string
	for {
		if err := cqClient.RateLimiter.Wait(ctx); err != nil {
			return nil
		}

		req := hubspotClient.BasicApi.
			GetPage(hubspot.WithAuthorizer(ctx, cqClient.Authorizer)).
			Properties(cqClient.Spec.TableOptions.ForTable("hubspot_crm_contacts").GetProperties()).
			Associations(cqClient.Spec.TableOptions.ForTable("hubspot_crm_contacts").GetAssociations()).
			Limit(client.MaxPageSize)

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

func syncContactsIncrementally(ctx context.Context, cqClient *client.Client, hubspotClient *contacts.APIClient, res chan<- any) error {
	const tableName = "hubspot_crm_contacts"

	lastModifiedDate, err := getLastModifiedDate(ctx, cqClient, tableName)
	if err != nil {
		return err
	}

	var after string
	for {
		if err := cqClient.RateLimiter.Wait(ctx); err != nil {
			return nil
		}

		req := contacts.PublicObjectSearchRequest{
			Limit:      client.MaxPageSize,
			Properties: []string{},
			Sorts:      sortAscByField("lastmodifieddate"),
		}

		if !lastModifiedDate.IsZero() {
			val := lastModifiedDate.Format(time.RFC3339)
			req.FilterGroups = []contacts.FilterGroup{
				{
					Filters: []contacts.Filter{
						{
							PropertyName: "lastmodifieddate",
							Operator:     "GTE",
							Value:        &val,
						},
					},
				},
			}
		}

		if len(after) > 0 {
			req.After = after
		}

		out, _, err := hubspotClient.SearchApi.
			Search(hubspot.WithAuthorizer(ctx, cqClient.Authorizer)).
			PublicObjectSearchRequest(req).
			Execute()

		if err != nil {
			return fmt.Errorf("failed to execute search req: %w", err)
		}
		res <- out.Results

		if !out.HasPaging() {
			break
		}

		paging := out.GetPaging()
		if !paging.HasNext() {
			break
		}

		after = paging.GetNext().After
		if after == "" {
			break
		}

		if after == client.SearchApiMaxPaginationItemCount {
			newLastModifiedDate := out.Results[len(out.Results)-1].Properties["lastmodifieddate"]
			t, err := time.Parse(time.RFC3339, newLastModifiedDate)
			if err != nil {
				return fmt.Errorf("failed to parse new last modified date: %w", err)
			}
			lastModifiedDate = t
			after = ""
		}
	}

	if !lastModifiedDate.IsZero() {
		return setLastModifiedDate(ctx, cqClient, tableName, lastModifiedDate)
	}

	return nil
}
