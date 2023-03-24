package client

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/shenzhencenter/google-ads-pb/clients"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"golang.org/x/exp/slices"
)

func (c *Client) initCustomers(ctx context.Context, client *clients.CustomerClient, spec *Spec) error {
	ctx = c.addDeveloperToken(ctx)

	ids := []string{spec.LoginCustomerID}
	if len(spec.LoginCustomerID) == 0 {
		accessibleIDs, err := getAccessibleIDs(ctx, client)
		if err != nil {
			return err
		}
		ids = accessibleIDs
	}

	customers := make(map[string][]string) // this will include duplicate data for now
	saved := make(map[string]struct{})     // we will save child node only once

	query := gaql.Query(new(resources.CustomerClient), nil) + "\nWHERE customer_client.level = 1" // 0 = self, 1 = direct child
	unprocessed := slices.Clone(ids)
	var id string
	for len(unprocessed) > 0 {
		id, unprocessed = unprocessed[0], unprocessed[1:]

		if _, ok := customers[id]; ok {
			// already processed, don't repeat
			continue
		}

		customers[id] = make([]string, 0)

		children, err := c.getChildren(ctx, id, query)
		if err != nil {
			return err
		}

		for _, child := range children {
			childID := strconv.FormatInt(*child.Id, 10)
			saved[childID] = struct{}{}
			customers[id] = append(customers[id], childID)

			if child.GetManager() {
				// even if we processed this child already, we still can check in the loop beginning
				unprocessed = append(unprocessed, childID)
			}
		}
	}

	// add orphaned resources from top
	for _, id := range ids {
		if _, ok := saved[id]; ok {
			continue
		}
		customers[id] = append(customers[id], id) // self-link
	}

	c.customers = dedup(limitTo(customers, spec.Customers))
	return nil
}

func getAccessibleIDs(ctx context.Context, client *clients.CustomerClient) ([]string, error) {
	req := &services.ListAccessibleCustomersRequest{}
	resp, err := client.ListAccessibleCustomers(ctx, req)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(resp.ResourceNames))
	for i, rn := range resp.ResourceNames {
		ids[i] = strings.TrimPrefix(rn, "customers/")
	}

	return ids, nil
}

func limitTo(customers map[string][]string, only []string) map[string][]string {
	if len(only) == 0 {
		return customers
	}

	res := make(map[string][]string)
	processed := make(map[string]struct{})

	unprocessed := slices.Clone(only)
	var curr string
	for len(unprocessed) > 0 {
		curr, unprocessed = unprocessed[0], unprocessed[1:]
		if _, ok := processed[curr]; ok {
			continue
		}
		processed[curr] = struct{}{}
		if v, ok := customers[curr]; ok {
			res[curr] = append(res[curr], v...)
			unprocessed = append(unprocessed, v...)
			continue
		}

		// if we got here we're sure that curr isn't a key in customers, so we just add it to the res
		for k, vv := range customers {
			if slices.Contains(vv, curr) {
				// we're OK to have duplicates as we'll clear the dups in dedup
				res[k] = append(res[k], curr)
			}
		}
	}
	return res
}

// dedup makes sure that each value is saved only once
func dedup(customers map[string][]string) map[string][]string {
	saw := make(map[string]struct{}, len(customers))
	res := make(map[string][]string, len(customers))

	for k, vv := range customers {
		for _, v := range vv {
			if _, ok := saw[v]; ok {
				continue
			}

			saw[v] = struct{}{}
			res[k] = append(res[k], v)
		}
	}

	return res
}

func (c *Client) getChildren(ctx context.Context, id, query string) ([]*resources.CustomerClient, error) {
	ctx = addLoginCustomerID(ctx, id)

	req := &services.SearchGoogleAdsStreamRequest{
		CustomerId:        id,
		Query:             query,
		SummaryRowSetting: enums.SummaryRowSettingEnum_NO_SUMMARY_ROW,
	}

	var res []*resources.CustomerClient
	ch := make(chan any)

	done := make(chan struct{})
	go func() {
		defer close(done)
		for v := range ch {
			res = append(res, v.(*resources.CustomerClient))
		}
	}()

	resp, err := c.GoogleAdsClient.SearchStream(ctx, req)
	if err != nil {
		return nil, err
	}

	err = ReceiveStream(resp.Recv, (*services.GoogleAdsRow).GetCustomerClient, ch)
	close(ch)
	if err != nil {
		return nil, err
	}

	<-done
	return res, nil
}
