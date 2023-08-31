package client

import (
	"context"
	"io"
)

func (c *Client) refreshIndex(ctx context.Context, index string) error {
	resp, err := c.typedClient.Indices.Refresh().Index(index).Do(ctx)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	return nil
}
