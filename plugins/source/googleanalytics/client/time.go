package client

import (
	"context"
	"time"
)

func genDates(ctx context.Context, c *Client, table string) (<-chan time.Time, error) {
	res, err := c.backend.GetKey(ctx, table+c.ID())
	if err != nil {
		return nil, err
	}

	startDate := c.StartDate
	if len(res) > 0 {
		startDate = res
	}

	// parse
	const dateLayout = "2006-01-02"
	t, err := time.Parse(dateLayout, startDate)
	if err != nil {
		return nil, err
	}

	ch := make(chan time.Time)
	go func() {
		defer close(ch)

		today := time.Now().UTC()

		for !t.After(today) {
			ch <- t
			t = t.Add(24 * time.Hour)
		}
	}()

	return ch, nil
}
