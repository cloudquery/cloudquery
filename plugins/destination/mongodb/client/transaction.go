package client

import (
	"context"
	"fmt"
)

func (c *Client) runWrite(ctx context.Context, op func(ctx context.Context) error) error {
	if c.spec == nil || c.spec.WriteRetry == nil || c.spec.WriteRetry.MaxAttempts <= 1 || !c.spec.WriteRetry.UseTransactions {
		return op(ctx)
	}
	session, err := c.client.StartSession()
	if err != nil {
		return fmt.Errorf("start session: %w", err)
	}
	defer session.EndSession(ctx)
	_, err = session.WithTransaction(ctx, func(sctx context.Context) (any, error) {
		return nil, op(sctx)
	})
	return err
}
