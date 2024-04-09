package crm

import (
	"context"
	"fmt"
	"github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/contacts"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"strconv"
	"time"
)

func fetchContacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := contacts.NewConfiguration()

	hubspotClient := contacts.NewAPIClient(config)
	cqClient := meta.(*client.Client)

	const tableName = "hubspot_crm_contacts"

	after := ""

	lastModifiedDate, err := getLastModifiedDate(ctx, cqClient, tableName)
	if err != nil {
		return err
	}

	for {
		if err := cqClient.RateLimiter.Wait(ctx); err != nil {
			return nil
		}

		req := contacts.PublicObjectSearchRequest{
			Limit:      client.SearchApiMaxPageSize,
			Properties: []string{},
			Sorts: []string{
				`{"propertyName": "lastmodifieddate", "direction": "ASCENDING}"`,
			},
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

		if after == strconv.Itoa(client.SearchApiMaxPaginationItemCount) {
			newLastModifiedDate := out.Results[len(out.Results)-1].Properties["lastmodifieddate"]
			t, err := time.Parse(time.RFC3339, newLastModifiedDate)
			if err != nil {
				return err
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
