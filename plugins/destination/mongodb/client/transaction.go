package client

import (
	"context"
	"fmt"
)

// runWrite executes op directly when transactions are disabled, or wraps it in
// session.WithTransaction when the spec opts in. The session is ended after the
// callback returns so connections are released back to the pool.
//
// Transactions require a replica set or sharded cluster (e.g. MongoDB Atlas).
// On standalone MongoDB the driver surfaces a clear error at write time.
func (c *Client) runWrite(ctx context.Context, op func(ctx context.Context) error) error {
	if c.spec == nil || c.spec.WriteRetry.MaxAttempts == 1 {
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
