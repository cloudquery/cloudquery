package incidents

import (
	"context"
	"time"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchIncidents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	// Time step incidents fetching in ±90 days intervals because the list incident API
	// won't return more than 10 000 incidents when using DateRange: "all".
	// There isn't any deep rational about the 90 days, it's just a number that should work for most cases
	// with an average of 111 incidents per day.
	timeStep := time.Hour * 24 * 90

	// Make an initial request to get the oldest incident.
	// The list incident endpoint combined with DateRange: "all" returns the incidents in ascending order of creation date.
	response, err := cqClient.PagerdutyClient.ListIncidentsWithContext(ctx, pagerduty.ListIncidentsOptions{
		Limit:     1,
		Offset:    0,
		DateRange: "all",
		TeamIDs:   cqClient.Spec.TeamIds,
	})
	if err != nil {
		return err
	}
	if len(response.Incidents) == 0 {
		return nil
	}

	since, err := time.Parse(time.RFC3339, response.Incidents[0].CreatedAt)
	if err != nil {
		return err
	}
	until := since.Add(timeStep)
	now := time.Now()

	for since.Before(now) {
		more := true
		offset := uint(0)
		for more {
			response, err := cqClient.PagerdutyClient.ListIncidentsWithContext(ctx, pagerduty.ListIncidentsOptions{
				Limit:   client.MaxPaginationLimit,
				Offset:  offset,
				Since:   since.UTC().Format(time.RFC3339), // since is inclusive
				Until:   until.UTC().Format(time.RFC3339), // until is exclusive
				TeamIDs: cqClient.Spec.TeamIds,
			})
			if err != nil {
				return err
			}
			if len(response.Incidents) == 0 {
				break
			}
			res <- response.Incidents
			offset += uint(len(response.Incidents))
			more = response.More
		}
		since = until
		until = since.Add(timeStep)
	}

	return nil
}
