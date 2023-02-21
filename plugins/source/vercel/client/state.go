package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
)

func (c *Client) GetPaginator(ctx context.Context, key string, ids ...string) (vercel.Paginator, error) {
	var pg vercel.Paginator

	if c.Backend == nil {
		return pg, nil
	}

	id := strings.Join(append([]string{c.ID()}, ids...), ":")
	value, err := c.Backend.Get(ctx, key, id)
	if err != nil {
		return pg, fmt.Errorf("failed to retrieve state from backend: %w", err)
	}
	if value == "" {
		return pg, nil
	}

	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return pg, fmt.Errorf("retrieved invalid state value: %q %w", value, err)
	}
	pg.Next = &val
	return pg, nil
}

func (c *Client) SavePaginator(ctx context.Context, key string, pg vercel.Paginator, ids ...string) error {
	if c.Backend == nil || pg.Next == nil || *pg.Next == 0 {
		return nil
	}

	id := strings.Join(append([]string{c.ID()}, ids...), ":")
	return c.Backend.Set(ctx, key, id, strconv.FormatInt(*pg.Next, 10))
}
