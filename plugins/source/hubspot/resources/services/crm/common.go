package crm

import (
	"context"
	"fmt"
	"github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/objects"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"time"
)

const (
	// maxPageSize is the largest page size that HubSpot allows.
	maxPageSize = 100

	// searchApiMaxPaginationItemCount is the pagination hard limit imposed by the Hubspot Search API. Stored as a string because that's the type being compared to.
	searchApiMaxPaginationItemCount = "10000"
)

func syncIncrementally(ctx context.Context, cqClient *client.Client, res chan<- any, objectType, lastModifiedDateFieldName string) error {
	hubspotClient := objects.NewAPIClient(objects.NewConfiguration())

	tableName := "hubspot_crm_" + objectType

	lastModifiedDate, err := getLastModifiedDate(ctx, cqClient, tableName)
	if err != nil {
		return err
	}

	var after string
	for {
		if err := cqClient.RateLimiter.Wait(ctx); err != nil {
			return nil
		}

		req := objects.PublicObjectSearchRequest{
			Limit:      maxPageSize,
			Properties: []string{},
			Sorts:      sortAscByField(lastModifiedDateFieldName),
		}

		if !lastModifiedDate.IsZero() {
			val := lastModifiedDate.Format(time.RFC3339)
			req.FilterGroups = []objects.FilterGroup{
				{
					Filters: []objects.Filter{
						{
							PropertyName: lastModifiedDateFieldName,
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
			Search(hubspot.WithAuthorizer(ctx, cqClient.Authorizer), objectType).
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

		if after == searchApiMaxPaginationItemCount && len(out.Results) > 0 {
			newLastModifiedDate := out.Results[len(out.Results)-1].Properties[lastModifiedDateFieldName]
			t, err := time.Parse(time.RFC3339, newLastModifiedDate)
			if err != nil {
				return fmt.Errorf("failed to parse new last modified date for table %s: %w", tableName, err)
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

func setLastModifiedDate(ctx context.Context, c *client.Client, tableName string, t time.Time) error {
	if err := c.Backend.SetKey(ctx, generateKey(c, tableName), t.Format(time.RFC3339Nano)); err != nil {
		return fmt.Errorf("failed to store cursor to backend: %w", err)
	}
	return nil
}

func getLastModifiedDate(ctx context.Context, c *client.Client, tableName string) (time.Time, error) {
	value, err := c.Backend.GetKey(ctx, generateKey(c, tableName))
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to retrieve state from backend: %w", err)
	}
	if value == "" {
		return time.Time{}, nil
	}
	t, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return time.Time{}, fmt.Errorf("retrieved invalid state value: %q %w", value, err)
	}
	return t, nil
}

func generateKey(c *client.Client, tableName string) string {
	return fmt.Sprintf("%s-%s-last-modified-date", c.ID(), tableName)
}

func sortAscByField(fieldName string) []string {
	return []string{
		// there is no generated type for this value, and it's a very simple JSON string, so we build it manually
		`{"propertyName":"` + fieldName + `","direction":"ASCENDING}"`,
	}
}
